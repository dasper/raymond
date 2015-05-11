package raymond

import "fmt"

// A partial template
type Partial struct {
	name   string
	source string
}

// All registered partials
var partials map[string]*Partial

func init() {
	partials = make(map[string]*Partial)
}

// Instanciate a new partial
func NewPartial(name string, source string) *Partial {
	return &Partial{
		name:   name,
		source: source,
	}
}

// Registers a new partial
func RegisterPartial(name string, source string) {
	if partials[name] != nil {
		panic(fmt.Errorf("Partial already registered: %s", name))
	}

	partials[name] = NewPartial(name, source)
}