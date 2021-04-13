// All extensions here are not part of the official GraphQL introspection specification.
// They are intended to be used via own operations using the introspection capabilities.

package extensions

import (
	"encoding/json"

	"github.com/graph-gophers/graphql-go/internal/common"
)

// NewDirectiveList creates a new *DirectiveList out of the given internal one
func NewDirectiveList(list common.DirectiveList) DirectiveList {
	if len(list) == 0 {
		return DirectiveList{list: make(map[string]DirectiveItem)}
	}

	directives := make(map[string]DirectiveItem, len(list))
	for _, d := range list {
		if len(d.Args) == 0 {
			directives[d.Name.Name] = DirectiveItem{
				Name: d.Name.Name,
			}

			continue
		}

		args := make(map[string]DirectiveArg, len(d.Args))
		for _, arg := range d.Args {
			args[arg.Name.Name] = DirectiveArg{
				Name:  arg.Name.Name,
				Value: arg.Value,
			}
		}
		directives[d.Name.Name] = DirectiveItem{
			Name: d.Name.Name,
			Args: DirectiveArgList{args},
		}

	}

	return DirectiveList{list: directives}
}

// DirectiveList is the outer API for operating with directives in the schema definition
type DirectiveList struct {
	list map[string]DirectiveItem
}

// Get returns a mapped directive by its given name.
// If the directive cannot be found (e.g. invalid or not attached) the method returns nil
func (l DirectiveList) Get(name string) *DirectiveItem {
	if d, ok := l.list[name]; ok {
		return &d
	}

	return nil
}

// List will return the attached DirectiveItems as an array
func (l DirectiveList) List() []DirectiveItem {
	r := make([]DirectiveItem, 0, len(l.list))
	for _, d := range l.list {
		r = append(r, d)
	}
	return r
}

// DirectiveItem represents a directive attached to an object, field, etc. directly in the schema
type DirectiveItem struct {
	Name string
	Args DirectiveArgList
}

type DirectiveArgList struct {
	args map[string]DirectiveArg
}

// Get returns a mapped directive argument by its given name.
// If the directive argument cannot be found (e.g. invalid or not attached) the method returns nil
func (l DirectiveArgList) Get(name string, dest interface{}) (_ bool, err error) {
	if arg, ok := l.args[name]; ok {
		var b []byte
		if b, err = json.Marshal(arg.Value.Value(nil)); err != nil {
			return false, err
		}

		if err = json.Unmarshal(b, dest); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

// DirectiveArg represents an argument provided to a directive attached to an element directly in the schema
type DirectiveArg struct {
	Name  string
	Value common.Literal
}
