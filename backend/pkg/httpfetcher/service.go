package httpfetcher

import (
	"io"
	"net/http"
)

func NewService() {
}

func fetch() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
}
