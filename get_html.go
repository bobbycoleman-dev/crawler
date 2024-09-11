package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("a network error occurred: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get URL: %v", res.Status)
	}

	header := res.Header.Get("content-type")
	if !strings.Contains(header, "text/html") {
		return "", fmt.Errorf("the URL is not an HTML page")
	}

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the response body: %v", err)
	}

	return string(html), nil
}
