package articleGql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/dyalicode/with-graphql-using-gpt/article/model"
)

// Create article type
var ArticleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.String},
			"title":   &graphql.Field{Type: graphql.String},
			"content": &graphql.Field{Type: graphql.String},
		},
	},
)

// Define article queries
var ArticleQueryFields = graphql.Fields{
	"articles": &graphql.Field{
		Type:    graphql.NewList(ArticleType),
		Resolve: GetArticlesResolver,
	},
	"article": &graphql.Field{
		Type:    ArticleType,
		Args:    graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.String}},
		Resolve: GetArticleByIDResolver,
	},
	"articleByTitle": &graphql.Field{
		Type:    graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "return article By Title", nil
		},
	},
	"articleByAuthor": &graphql.Field{
		Type:    graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "return article By Author", nil
		},
	},
}

// Create a separate object to hold the article queries
var ArticleQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "ArticleQuery",
	Fields: ArticleQueryFields,
})

// Resolver functions for article queries
func GetArticlesResolver(p graphql.ResolveParams) (interface{}, error) {
	// Implement your logic here
	return []articleModel.Article{
		{ID: "1", Title: "Article 1", Content: "Content 1"},
		{ID: "2", Title: "Article 2", Content: "Content 2"},
	}, nil
}

func GetArticleByIDResolver(p graphql.ResolveParams) (interface{}, error) {
	// Implement your logic here
	articleID, ok := p.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("Invalid article ID")
	}

	// Query the database or perform any required operations to get the article by ID
	// For now, returning a dummy article
	return articleModel.Article{ID: articleID, Title: "Sample Article", Content: "Sample Content"}, nil
}
