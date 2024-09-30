package kafka

type Topic string

const (
	TopicClimateMeasurements Topic = "climate-measurements"
	TopicSeederMessages      Topic = "seeder-messages"
)

var Topics = []Topic{
	TopicClimateMeasurements, TopicSeederMessages,
}
