package expectations

// Expectation defines a pluggable interface for matching desired state to actual
type Expectation interface {
	Matches(string) bool
	String() string
}

// Params defines the parameters that construct an Expectation
type Params struct {
	Type string
	Data map[string]string
}

// New creates a new Expectation from the given Params
func New(p Params) Expectation {
	item, ok := types[p.Type]
	if !ok {
		item = types["error"]
		p.Data["type"] = p.Type
	}
	return item(p.Data)
}

type builder func(map[string]string) Expectation

var types = map[string]builder{
	"error":  newErrorExpectation,
	"string": newStringExpectation,
	"regex":  newRegexExpectation,
	"set":    newSetExpectation,
}
