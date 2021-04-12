// All extensions here are not part of the official GraphQL introspection specification.
// They are intended to be used via own operations using the introspection capabilities.

package extensions

import "github.com/graph-gophers/graphql-go/internal/common"

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

		args := make([]DirectiveArg, len(d.Args))
		for j, a := range d.Args {
			args[j] = DirectiveArg{
				Name:  a.Name.Name,
				Value: a.Value.String(),
			}
		}
		directives[d.Name.Name] = DirectiveItem{
			Name: d.Name.Name,
			Args: &args,
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
	Args *[]DirectiveArg
}

// DirectiveArg sdf
type DirectiveArg struct {
	Name  string
	Value string
}
