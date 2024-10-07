package climateMeasurements

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/configs"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka/dtos"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"

	rskafka "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	climateMeasurementLatestPath  = "/cgi-bin/tou3.cgi"
	climateMeasurementHistoryPath = "cgi-bin/temperaturen"
	args                          = uint32(10)
	max                           = uint32(999999)
)

type Fetcher struct {
	KafkaWriter *kafka.Writer
	Config      configs.FetcherConfig
}

func NewFetcher(configFetcher configs.FetcherConfig, configKafka rskafka.Config) (*Fetcher, error) {
	writer, err := rskafka.InitWriter(configKafka, rskafka.TopicClimateMeasurements)
	if err != nil {
		return nil, fmt.Errorf("Kafka connection error %v", err)
	}

	return &Fetcher{Config: configFetcher, KafkaWriter: writer}, nil
}

func (f *Fetcher) Close() {
	defer f.KafkaWriter.Close()
}

func (f *Fetcher) FetchLatest(count uint32) bool {
	if count == 0 {
		count = args
	}
	count = min(count, max)
	url, _ := url.JoinPath(f.Config.Url, climateMeasurementLatestPath)
	url = fmt.Sprintf("%s?%d", url, count)

	measurements, err := fetch(url)
	if err != nil {
		logging.Logger.Error("Error retrieving data from url", zap.String("url", url), zap.Error(err))
		return false
	}

	err = f.send(measurements)
	if err != nil {
		logging.Logger.Error("Send error", zap.Error(err))
		return false
	}

	return true
}

func (f *Fetcher) FetchHistory(rooms []models.Room) {
	for _, room := range rooms {

		url, _ := url.JoinPath(f.Config.Url, climateMeasurementHistoryPath, room.ExternalId)

		logging.Logger.Info("Opening url %s for room %s", zap.String("URL", url), zap.String("room", room.ExternalId))

		measurements, err := fetch(url)
		if err != nil {
			logging.Logger.Error("Error retrieving data from url", zap.String("url", url), zap.Error(err))
			continue
		}

		logging.Logger.Info("Got measurements cnt %d", zap.Int("CNT", len(measurements)))

		err = f.send(measurements)
		if err != nil {
			logging.Logger.Error("Send error", zap.Error(err))
		}
	}
}

func fetch(url string) ([]dtos.ClimateMeasurement, error) {
	measurements := []dtos.ClimateMeasurement{}

	logging.Logger.Sugar().Debugf("CALLING URL %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return measurements, fmt.Errorf("HTTP error: %v", err)
	}

	if resp.StatusCode != 200 {
		return measurements, fmt.Errorf("HTTP status invalid %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return measurements, err
	}

	// log.Printf("body %v", string(body))
	measurements, err = parseLatest(string(body))
	if err != nil {
		return measurements, fmt.Errorf("Parse error: %v", err)
	}

	return measurements, nil
}

func parseLatest(value string) ([]dtos.ClimateMeasurement, error) {
	measurements := []dtos.ClimateMeasurement{}
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
			timestamp := parts[len(parts)-2] + " " + parts[len(parts)-1]

			if err1 == nil && err2 == nil {
				measurements = append(measurements, dtos.ClimateMeasurement{RoomId: currentHeader, Temperature: float32(temperature), Humidity: float32(humidity), Timestamp: timestamp})
			}
		}
	}

	return measurements, nil
}

func (f *Fetcher) send(measurements []dtos.ClimateMeasurement) error {
	if len(measurements) <= 0 {
		return nil
	}

	msgs := make([]kafka.Message, len(measurements))

	logging.Logger.Sugar().Debugf("Sending to kafka cnt %d", len(msgs))

	for idx, measurement := range measurements {
		jsonData, err := json.Marshal(measurement)
		if err != nil {
			return fmt.Errorf("Error marshaling data %v", err)
		}

		msgs[idx] = kafka.Message{
			Key:   []byte(measurement.RoomId),
			Value: jsonData,
		}
	}

	f.KafkaWriter.WriteMessages(context.Background(), msgs...)

	return nil
}
