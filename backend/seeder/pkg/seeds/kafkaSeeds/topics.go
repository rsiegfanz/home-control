package kafkaSeeds

import (
	"fmt"
	"net"
	"strconv"

	rskafka "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"

	"github.com/segmentio/kafka-go"
)

func SeedTopics(config rskafka.Config) error {
	conn, err := kafka.Dial("tcp", config.Host)
	if err != nil {
		return fmt.Errorf("Kafka connection error %v", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("Kafka controller error %v", err)
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return fmt.Errorf("Kafka dial error %v", err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{}

	for _, topic := range rskafka.Topics {
		topicConfigs = append(topicConfigs, kafka.TopicConfig{
			Topic:             string(topic),
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return fmt.Errorf("Create topics error %v", err)
	}

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return fmt.Errorf("Read partitions error %v", err)
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	logging.Logger.Sugar().Infof("Kafka partitions", m)

	return nil
}
