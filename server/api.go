package main

import (
	"encoding/json"
	"net/http"
	"database/sql"
)

func MetricsReceiver(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m Metric
		json.NewDecoder(r.Body).Decode(&m)
		InsertMetric(db, m)
		w.WriteHeader(http.StatusOK)
	}
}

func MetricsList(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, _ := db.Query("SELECT node,cpu,memory,disk,load,network_in,network_out,timestamp FROM metrics ORDER BY timestamp DESC LIMIT 100")
		var metrics []Metric
		for rows.Next() {
			var m Metric
			rows.Scan(&m.Node,&m.CPU,&m.Memory,&m.Disk,&m.Load,&m.NetworkIn,&m.NetworkOut,&m.Timestamp)
			metrics = append(metrics,m)
		}
		json.NewEncoder(w).Encode(metrics)
	}
}
