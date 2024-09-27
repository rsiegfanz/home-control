package kafka

import "github.com/segmentio/kafka-go"

func InitWriter(config Config, topic string) (*kafka.Writer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{config.Host},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	// defer writer.Close()
	return writer, nil
}
