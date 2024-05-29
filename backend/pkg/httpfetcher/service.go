package httpfetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rs/homecontrol/pkg/models"
)

func RunService() {
	go fetchLatest()
}

func fetchLatest() {
	for {
		sleep()

		log.Println("fetch latest")

		measurements, err := fetch("", "")
		if err != nil {
			log.Println("fetch error: ", err)
			continue
		}

	}
}

func fetch(url string, savePath string) ([]models.Measurement, error) {
	measurements := []models.Measurement{}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("http error: ", err)
		return measurements, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	measurements, err = parseLatest(string(body))
	if err != nil {
		log.Println("parse error: ", err)
		return measurements, err
	}

	return measurements, nil
}

func parseLatest(value string) ([]models.Measurement, error) {
	measurements := []models.Measurement{}
	if value == "" {
		return measurements, fmt.Errorf("no data received")
	}

	return measurements, nil
}

func sleep() {
	time.Sleep(5 * time.Second)
}
