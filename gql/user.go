package gql

import (
	"fmt"
	"github.com/dyalicode/with-graphql-using-gpt/models"
	"github.com/graphql-go/graphql"
)

// Define user queries
var UserQueryFields = graphql.Fields{
	"users": &graphql.Field{
		Type:    graphql.NewList(UserType),
		Resolve: GetUsersResolver,
	},
	"user": &graphql.Field{
		Type:    UserType,
		Args:    graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.String}},
		Resolve: GetUserByIDResolver,
	},
}

// Create user type
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"name":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
		},
	},
)

// Create a separate object to hold the user queries
var UserQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "UserQuery",
	Fields: UserQueryFields,
})

// Resolver functions for user queries
func GetUsersResolver(p graphql.ResolveParams) (interface{}, error) {
	// Implement your logic here
	return []models.User{
		{ID: "1", Name: "User 1", Email: "user1@example.com"},
		{ID: "2", Name: "User 2", Email: "user2@example.com"},
	}, nil
}

func GetUserByIDResolver(p graphql.ResolveParams) (interface{}, error) {
	// Implement your logic here
	userID, ok := p.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("Invalid user ID")
	}

	// Query the database or perform any required operations to get the user by ID
	// For now, returning a dummy user
	return models.User{ID: userID, Name: "Sample User", Email: "sampleuser@example.com"}, nil
}
