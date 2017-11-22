package queries

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields return all the available queries.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"user": GetUserQuery(),
	}
}