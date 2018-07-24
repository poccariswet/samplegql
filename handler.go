package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type user struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/*
   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						u := user{}
						if err := DB.Get(&u, "SELECT * FROM sampleql WHERE id=$1", idQuery); err != nil {
							return nil, nil
						}
						return &u, nil
					}
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	res := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(res.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", res.Errors)
	}

	return res
}

func QLhandler(w http.ResponseWriter, r *http.Request) {
	res := executeQuery(r.URL.Query().Get("query"), schema) //ex){user(id:"1"){name}}
	json.NewEncoder(w).Encode(res)
}
