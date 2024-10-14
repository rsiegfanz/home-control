package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/graphql-go/graphql"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/resolvers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/schemas"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

	subscriptionFields := graphql.Fields{
		"climateMeasurementUpdates": &graphql.Field{
			Type: schemas.ClimateMeasurementSchema,
			Args: graphql.FieldConfigArgument{
				"roomExternalId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: queryResolver.SubscribeToClimateMeasurements,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootSubscription := graphql.ObjectConfig{Name: "Subscription", Fields: subscriptionFields}

	schemaConfig := graphql.SchemaConfig{
		Query:        graphql.NewObject(rootQuery),
		Subscription: graphql.NewObject(rootSubscription),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		logging.Logger.Panic("invalid schema", zap.Error(err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if websocket.IsWebSocketUpgrade(r) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				logging.Logger.Error("WebSocket upgrade failed", zap.Error(err))
				return
			}
			defer conn.Close()

			HandleSubscription(conn, &schema, queryResolver)
			return
		}

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

func HandleSubscription(conn *websocket.Conn, schema *graphql.Schema, resolver *resolvers.QueryResolver) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var params struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal(message, &params); err != nil {
			continue
		}

		go func() {
			result := graphql.Do(graphql.Params{
				Schema:        *schema,
				RequestString: params.Query,
				Context:       context.Background(),
			})

			if result.HasErrors() {
				conn.WriteJSON(map[string]interface{}{
					"errors": result.Errors,
				})
				return
			}

			if data, ok := result.Data.(map[string]interface{}); ok {
				if stream, ok := data["climateMeasurementUpdates"].(chan interface{}); ok {
					for update := range stream {
						conn.WriteJSON(map[string]interface{}{
							"data": map[string]interface{}{
								"climateMeasurementUpdates": update,
							},
						})
					}
				}
			}
		}()
	}
}
