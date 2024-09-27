package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func InitWriter(config Config, topic string) (*kafka.Writer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{config.Host},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	data, _ := json.Marshal("data")
	error := writer.WriteMessages(context.Background(), kafka.Message{Key: []byte("test"), Value: data})
	log.Printf("kafka write result %v", error)
	// defer writer.Close()
	return writer, nil
}

/*
func InitReader() {
reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{config.KafkaURL},
        GroupID:   "worker-group",
        Topic:     "sensor-data",
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })
}*/
