package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "host=localhost port=5432 user=postgres password=pass dbname=monitor sslmode=disable")
	InitDB(db)

	http.HandleFunc("/metrics", MetricsReceiver(db))
	http.HandleFunc("/api/metrics", MetricsList(db))
	http.HandleFunc("/ws", WSHandler)
	http.Handle("/", http.FileServer(http.Dir("../web")))

	fmt.Println("Monitoring server running on :8080")
	http.ListenAndServe(":8080", nil)
}
