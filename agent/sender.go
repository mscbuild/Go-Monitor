package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Send(serverURL string, m Metrics) error {
	body, _ := json.Marshal(m)
	_, err := http.Post(serverURL, "application/json", bytes.NewBuffer(body))
	return err
}
