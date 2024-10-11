package schemas

import "github.com/graphql-go/graphql"

var ClimateMeasurementType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClimateMeasurement",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"recordedAt": &graphql.Field{
				Type: graphql.String,
			},
			"roomExternalId": &graphql.Field{
				Type: graphql.String,
			},
			"temperature": &graphql.Field{
				Type: graphql.Float,
			},
			"humidity": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
