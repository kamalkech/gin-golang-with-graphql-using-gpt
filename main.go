package main

import (
	// "github.com/dyalicode/with-graphql-using-gpt/gql"
	// "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	// "log"
	"net/http"
)

func main() {

	// Get schema generated
	schema, _ := UseSchema()

	// Create a GraphQL HTTP handler with the schema
	graphQLHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Serve the GraphQL API over HTTP
	http.Handle("/graphql", graphQLHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandboxHTML)
	}))

	http.ListenAndServe(":8080", nil)
}

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: "http://localhost:8080/graphql",
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>
</body>

</html>`)
