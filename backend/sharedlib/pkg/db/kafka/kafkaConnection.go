package kafka

import (
	"github.com/segmentio/kafka-go"
)

var groupId = "rs"

func InitWriter(config Config, topic Topic) (*kafka.Writer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:    []string{config.Host},
		Topic:      string(topic),
		Balancer:   &kafka.LeastBytes{},
		BatchSize:  1000,
		BatchBytes: 10e6,
	})

	// defer writer.Close()
	return writer, nil
}

func InitReader(config Config, topic Topic) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.Host},
		GroupID:   groupId,
		Partition: 0,
		Topic:     string(topic),
		//		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	// defer reader.Close()
	return reader
}
