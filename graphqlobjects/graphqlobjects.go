package graphqlobjects

import (
  "github.com/graphql-go/graphql"

  "github.com/ccleung/playground/dataobjects"
)

var(
  Organization *graphql.Object
  User *graphql.Object
)

func init() {
   User = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
     "id": &graphql.Field{
       Type: graphql.String,
     },
     "name": &graphql.Field{
       Type: graphql.String,
     },
   },
  })
  
  Organization = graphql.NewObject(graphql.ObjectConfig{
    Name: "Organization",
    Fields: graphql.Fields{
      "user": &graphql.Field{
        Type: User,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return p.Source.(*dataobjects.Organization).User, nil
        },
      },
    },
  })
}
