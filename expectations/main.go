package expectations

import (
	"fmt"
)

// Expectation defines a pluggable interface for matching desired state to actual
type Expectation interface {
	Load(map[string]string) Expectation
	Matches(string) bool
	String() string
}

// Wrapper defines the parameters that construct an Expectation
type Wrapper struct {
	Type        string            `json:"type"`
	Data        map[string]string `json:"data"`
	expectation Expectation
}

// Matches proxies the request to the underlying expectation
func (w Wrapper) Matches(actual string) bool {
	if w.expectation == nil {
		w.load()
	}
	return w.expectation.Matches(actual)
}

// String proxies the request to the underlying expectation
func (w Wrapper) String() string {
	if w.expectation == nil {
		w.load()
	}
	return w.expectation.String()
}

func (w Wrapper) load() {
	itemFunc, ok := types[w.Type]
	if !ok {
		itemFunc = types["error"]
		w.Data["msg"] = fmt.Sprintf("expectation type not known: %s", w.Type)
	}
	e := itemFunc()
	w.expectation = e.Load(w.Data)
}

type builder func() Expectation

var types = map[string]builder{
	"error":  func() Expectation { return &errorExpectation{} },
	"string": func() Expectation { return &stringExpectation{} },
	"regex":  func() Expectation { return &regexExpectation{} },
	"set":    func() Expectation { return &setExpectation{} },
}
