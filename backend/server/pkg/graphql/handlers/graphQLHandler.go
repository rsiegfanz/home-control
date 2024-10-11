package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/resolvers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/schemas"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func NewGraphQLHandler(queryResolver *resolvers.QueryResolver) http.HandlerFunc {
	fields := graphql.Fields{
		"climateMeasurements": &graphql.Field{
			Type:        graphql.NewList(schemas.ClimateMeasurementSchema),
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
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		logging.Logger.Panic("invalid schema", zap.Error(err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			VariableValues: params.Variables,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
