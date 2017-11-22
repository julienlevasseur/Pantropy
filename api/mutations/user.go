package mutations

import (
	"../types"

	"log"
	"github.com/graphql-go/graphql"
)

// GetCreateUserMutation creates a new user and returns it.
func GetCreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("[mutation] create user\n")

			user := &types.User{
				Firstname: params.Args["firstname"].(string),
				Lastname: params.Args["lastname"].(string),
			}

			return user, nil
		},
	}
}
