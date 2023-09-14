package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var apiCache = NewCache(time.Duration(300) * time.Second)

func Request(url string) (responseBody []byte, err error) {
	cachedResp, err := apiCache.Get(url)
	if len(cachedResp) != 0 {
		return cachedResp, nil
	}
	var empty []byte
	resp, err := http.Get(url)
	if err != nil {
		return empty, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return empty, errors.New(fmt.Sprintf("Got incorrect response %v from server", resp.StatusCode))
	}
	if err != nil {
		return empty, err
	}
	apiCache.Add(url, body)
	return body, nil
}
