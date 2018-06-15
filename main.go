package main

import (
	"log"
	"net/http"
	// "runtime"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	// "github.com/graphql-go/relay/examples/starwars"
	"github.com/ccleung/playground/dataobjects"
	"github.com/ccleung/playground/graphqlobjects"
)

var userData = &dataobjects.User{ID: "1", Name: "Clement"}
var organizationData = &dataobjects.Organization{User: userData}

func main() {
	// custom schema
	fields := graphql.Fields{
		"Organization": &graphql.Field{
			Type: graphqlobjects.Organization,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return organizationData, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	// runtime.Breakpoint()

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// simplest relay-compliant graphql server HTTP handler
	// using Starwars schema from `graphql-relay-go` examples
	// schema = starwars.Schema

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/graphql", h)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
