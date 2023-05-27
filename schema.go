package main

import (
	"github.com/dyalicode/with-graphql-using-gpt/article/gql"
	"github.com/dyalicode/with-graphql-using-gpt/user/gql"
	"github.com/graphql-go/graphql"
	"log"
)

func UseSchema() (graphql.Schema, error){
	// Merge the article and user queries in the root query
	RootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: graphql.Fields{},
	})

	// ********************* Merge All queries. ************************//
	RootQuery.AddFieldConfig("queryArticle", &graphql.Field{
		Type: articleGql.ArticleQuery,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	})
	RootQuery.AddFieldConfig("queryUser", &graphql.Field{
		Type: userGql.UserQuery,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "query user", nil
		},
	})

	// ********************* Merge All mutation. ************************//

	// Create the schema using the root query
	SchemaConfig := graphql.SchemaConfig{
		Query: RootQuery,
	}
	Schema, err := graphql.NewSchema(SchemaConfig)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	return Schema, nil
}
