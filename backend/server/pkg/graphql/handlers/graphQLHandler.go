package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/resolvers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/schemas"
)

func NewGraphQLHandler(queryResolver *resolvers.QueryResolver) http.HandlerFunc {
	fields := graphql.Fields{
		"climateData": &graphql.Field{
			Type:        graphql.NewList(schemas.ClimateMeasurementType),
			Description: "Get climate data for a specific room and time period",
			Args: graphql.FieldConfigArgument{
				"startDate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"endDate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"roomExternalId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				startDate := p.Args["startDate"].(string)
				endDate := p.Args["endDate"].(string)
				roomExternalId := p.Args["roomExternalId"].(string)
				return queryResolver.GetClimateMeasurements(p.Context, startDate, endDate, roomExternalId)
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)

	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query string `json:"query"`
		}
		_ = json.NewDecoder(r.Body).Decode(&params)
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})
		json.NewEncoder(w).Encode(result)
	}
}
