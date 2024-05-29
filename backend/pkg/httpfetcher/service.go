package httpfetcher

import (
	"io"
	"log"
	"net/http"
)

func NewService() {
}

func fetch(url string, savePath: string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http error: ", err))
    return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
}
