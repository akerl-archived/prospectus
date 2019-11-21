package plugin

const (
	prospectusDirName = ".prospectus.d"
)

// LoadInput defines the input passed to a plugin to load checks
type LoadInput struct {
	Dir  string `json:"dir"`
	File string `json:"file"`
}

func (l LoadInput) Load() CheckSet {
	cs := CheckSet{}
	err := call(input.File, "load", input, &cs)
	if err != nil {
		cs = CheckSet{Check{Name: "__failure_to_load__"}}
	}
	for index := range cs {
		cs[index].Dir = input.Dir
		cs[index].File = input.File
	}
	return cs
}

// NewSet returns a CheckSet based on a provided list of directories
func NewSet(relativeDirs []string) (CheckSet, error) {
	var err error

	dirs := make([]string, len(relativeDirs))
	for index, item := range relativeDirs {
		dirs[index], err = filepath.Abs(item)
		if err != nil {
			return CheckSet{}, err
		}
	}

	var cs CheckSet
	for _, item := range dirs {
		newSet, err := newSetFromDir(item)
		if err != nil {
			return CheckSet{}, err
		}
		cs = append(cs, newSet...)
	}

	return cs, nil
}

func newSetFromDir(absoluteDir string) (CheckSet, error) {
	prospectusDir := filepath.Join(absoluteDir, prospectusDirName)

	fileObjs, err := ioutil.ReadDir(prospectusDir)
	if err != nil {
		return CheckSet{}, err
	}

	var cs CheckSet
	for _, fileObj := range fileObjs {
		file := filepath.Join(prospectusDir, fileObj.Name())
		newSet, err := newSetFromFile(absoluteDir, file)
		if err != nil {
			return CheckSet{}, err
		}
		cs = append(cs, newSet...)
	}

	return cs, nil
}

func newSetFromFile(dir, file string) (CheckSet, error) {
	cs := CheckSet{}
	input := loadCheckInput{Dir: dir}
	err := execProspectusFile(file, "load", input, &cs)
	if err != nil {
		return CheckSet{}, fmt.Errorf("Failed loading %s: %s", file, err)
	}
	for index := range cs {
		cs[index].Dir = dir
		cs[index].File = file
	}
	return cs, nil
}
