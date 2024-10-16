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
		logging.Logger.Debug("Checking WebSocket origin", zap.String("origin", r.Header.Get("Origin")))
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

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemas.ClimateMeasurementSubscription.Resolve = queryResolver.SubscribeToClimateMeasurements

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name: "Subscription",
			Fields: graphql.Fields{
				"climateMeasurementUpdates": schemas.ClimateMeasurementSubscription,
			},
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		logging.Logger.Panic("invalid schema", zap.Error(err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if websocket.IsWebSocketUpgrade(r) {
			logging.Logger.Info("WebSocket upgrade requested")
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				logging.Logger.Error("WebSocket upgrade failed", zap.Error(err))
				return
			}
			defer conn.Close()

			logging.Logger.Info("WebSocket connection established")
			HandleSubscription(conn, &schema, queryResolver)
			return
		}

		logging.Logger.Debug("Handling GraphQL HTTP request")

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

		logging.Logger.Debug("Executing GraphQL query",
			zap.String("operationName", params.OperationName),
			zap.Any("variables", params.Variables))

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
	logging.Logger.Info("Starting subscription handler")
	defer logging.Logger.Info("Subscription handler ended")

	subscriptions := make(map[string]context.CancelFunc)
	defer func() {
		for _, cancel := range subscriptions {
			cancel()
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logging.Logger.Error("Error reading WebSocket message", zap.Error(err))
			break
		}

		logging.Logger.Debug("Received WebSocket message", zap.ByteString("message", message))

		var operationMessage struct {
			Type    string          `json:"type"`
			ID      string          `json:"id,omitempty"`
			Payload json.RawMessage `json:"payload,omitempty"`
		}

		if err := json.Unmarshal(message, &operationMessage); err != nil {
			logging.Logger.Error("Failed to unmarshal message", zap.Error(err))
			continue
		}

		switch operationMessage.Type {
		case "connection_init":
			conn.WriteJSON(map[string]string{"type": "connection_ack"})
		case "subscribe":
			var params struct {
				Query     string                 `json:"query"`
				Variables map[string]interface{} `json:"variables"`
			}
			if err := json.Unmarshal(operationMessage.Payload, &params); err != nil {
				logging.Logger.Error("Failed to unmarshal subscription payload", zap.Error(err))
				continue
			}

			ctx, cancel := context.WithCancel(context.Background())
			subscriptions[operationMessage.ID] = cancel

			go func() {
				defer cancel()
				for {
					result := graphql.Do(graphql.Params{
						Schema:         *schema,
						RequestString:  params.Query,
						VariableValues: params.Variables,
						Context:        ctx,
					})

					if result.HasErrors() {
						logging.Logger.Error("GraphQL execution errors")
						conn.WriteJSON(map[string]interface{}{
							"type":    "error",
							"id":      operationMessage.ID,
							"payload": result.Errors,
						})
						return
					}

					if data, ok := result.Data.(map[string]interface{}); ok {
						err := conn.WriteJSON(map[string]interface{}{
							"type": "next",
							"id":   operationMessage.ID,
							"payload": map[string]interface{}{
								"data": data,
							},
						})
						if err != nil {
							logging.Logger.Error("Failed to write WebSocket message", zap.Error(err))
							return
						}
					} else {
						logging.Logger.Error("Unexpected data format in GraphQL result")
					}

					select {
					case <-ctx.Done():
						return
					default:
						// Continue to next iteration
					}
				}
			}()
		case "complete":
			logging.Logger.Info("Client requested to complete subscription", zap.String("id", operationMessage.ID))
			if cancel, ok := subscriptions[operationMessage.ID]; ok {
				cancel()
				delete(subscriptions, operationMessage.ID)
				conn.WriteJSON(map[string]interface{}{
					"type": "complete",
					"id":   operationMessage.ID,
				})
			}
		}
	}
}
