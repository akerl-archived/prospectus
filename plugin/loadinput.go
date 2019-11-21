package plugin

import (
	"io/ioutil"
	"path/filepath"
)

const (
	prospectusDirName = ".prospectus.d"
)

// LoadInput defines the input passed to a plugin to load checks
type LoadInput struct {
	Dir  string `json:"dir"`
	File string `json:"file"`
}

func (l LoadInput) Load() AttributeSet {
	cs := AttributeSet{}
	err := call(l.File, "load", l, &cs)
	if err != nil {
		cs = AttributeSet{Attribute{Name: "__failure_to_load__"}}
	}
	for index := range cs {
		cs[index].Dir = l.Dir
		cs[index].File = l.File
	}
	return cs
}

// NewSet returns a AttributeSet based on a provided list of directories
func NewSet(relativeDirs []string) (AttributeSet, error) {
	var err error

	dirs := make([]string, len(relativeDirs))
	for index, item := range relativeDirs {
		dirs[index], err = filepath.Abs(item)
		if err != nil {
			return AttributeSet{}, err
		}
	}

	as := AttributeSet{}
	for _, item := range dirs {
		newSet, err := newSetFromDir(item)
		if err != nil {
			return AttributeSet{}, err
		}
		as = append(as, newSet...)
	}

	return as, nil
}

func newSetFromDir(absoluteDir string) (AttributeSet, error) {
	prospectusDir := filepath.Join(absoluteDir, prospectusDirName)

	fileObjs, err := ioutil.ReadDir(prospectusDir)
	if err != nil {
		return AttributeSet{}, err
	}

	var as AttributeSet
	for _, fileObj := range fileObjs {
		file := filepath.Join(prospectusDir, fileObj.Name())
		newSet := newSetFromFile(absoluteDir, file)
		as = append(as, newSet...)
	}

	return as, nil
}

func newSetFromFile(dir, file string) AttributeSet {
	input := LoadInput{Dir: dir, File: file}
	return input.Load()
}
