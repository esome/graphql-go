package introspection

import "github.com/graph-gophers/graphql-go/introspection/extensions"

// Disclaimer: All extensions here are not part of the official spec. They are intended to be used via own operations
// using the introspection capabilities

// Directives provides a list of the directives attached to the field in the schema definition
func (r *Field) Directives() extensions.DirectiveList {
	return extensions.NewDirectiveList(r.field.Directives)
}
