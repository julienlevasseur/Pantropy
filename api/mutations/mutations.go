package mutations

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields return all the available mutations.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createUser": GetCreateUserMutation(),
	}
}
