package rooms

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/configs"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"

	rskafka "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type RoomFetcher struct {
	KafkaWriter *kafka.Writer
	Config      configs.FetcherConfig
}

func NewRoomFetcher(configFetcher configs.FetcherConfig, configKafka rskafka.Config) (*RoomFetcher, error) {
	writer, err := rskafka.InitWriter(configKafka, "rooms")
	if err != nil {
		return nil, fmt.Errorf("Kafka connection error", err)
	}

	return &RoomFetcher{Config: configFetcher, KafkaWriter: writer}, nil
}

func (f *RoomFetcher) Fetch(rooms []models.Room) {
	for _, room := range rooms {
		url := path.Join(f.Config.Url, room.ExternalId)
		measurements, err := fetch(url)
		if err != nil {
			logging.Logger.Error("Error retrieving data from url", zap.String("url", url), zap.Error(err))
			continue
		}

		err = f.send(measurements)
		if err != nil {
			logging.Logger.Error("Send error", zap.Error(err))
		}
	}
}

func fetch(url string) ([]MeasurementDto, error) {
	measurements := []MeasurementDto{}

	resp, err := http.Get(url)
	if err != nil {
		return measurements, fmt.Errorf("HTTP error: ", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return measurements, err
	}

	measurements, err = parseLatest(string(body))
	if err != nil {
		return measurements, fmt.Errorf("Parse error: ", err)
	}

	return measurements, nil
}

func parseLatest(value string) ([]MeasurementDto, error) {
	measurements := []MeasurementDto{}
	currentHeader := ""

	if strings.TrimSpace(value) == "" {
		return measurements, fmt.Errorf("no data received")
	}

	scanner := bufio.NewScanner(strings.NewReader(value))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "==" {
			currentHeader = ""
			continue
		}

		if strings.HasSuffix(line, ".werte") || line == "funkbme280" {
			currentHeader = line
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			temperature, err1 := strconv.ParseFloat(parts[0], 64)
			humidity, err2 := strconv.ParseFloat(parts[1], 64)

			if err1 == nil && err2 == nil {
				measurements = append(measurements, MeasurementDto{FileId: currentHeader, Temperature: float32(temperature), Humidity: float32(humidity)})
			}

		}
	}

	return measurements, nil
}

func (f *RoomFetcher) send(measurements []MeasurementDto) error {
	jsonData, err := json.Marshal(measurements)
	if err != nil {
		return fmt.Errorf("Error marshaling data", err)
	}

	for _, measurement := range measurements {
		// Schreibe die Daten zu Kafka
		err = f.KafkaWriter.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(measurement.FileId),
				Value: jsonData,
			},
		)
		if err != nil {
			return fmt.Errorf("Error writing to kafka", err)
		}
	}

	return nil
}
