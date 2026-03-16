package main

import (
	"fmt"
	"time"
)

func main() {
	serverURL := "http://localhost:8080/metrics"
	nodeName := "web-server-1"
	interval := 10

	fmt.Println("Agent started for node:", nodeName)

	for {
		m := Collect(nodeName)
		if err := Send(serverURL, m); err != nil {
			fmt.Println("Error sending metrics:", err)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
