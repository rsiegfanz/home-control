package httpfetcher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/homecontrol/pkg/repository"
)

func RunService(latestUrl string, latestFilePath string) {
	go fetchLatest(latestUrl, latestFilePath)
}

func fetchLatest(url string, filePath string) {
	for {
		sleep()

		log.Println("fetch latest")

		measurementDtos, err := fetch(url)
		if err != nil {
			continue
		}

		measurements := MapMeasurementDtosToModels(measurementDtos)

		repository.SaveLatestAll(filePath, measurements)
	}
}

func fetch(url string) ([]MeasurementDto, error) {
	measurements := []MeasurementDto{}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("http error: ", err)
		return measurements, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return measurements, err
	}

	measurements, err = parseLatest(string(body))
	if err != nil {
		log.Println("parse error: ", err)
		return measurements, err
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

		if strings.HasSuffix(line, ".werte") {
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

func sleep() {
	time.Sleep(5 * time.Second)
}
