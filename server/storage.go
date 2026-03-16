package main

import (
	"database/sql"
	"time"
)

func InitDB(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS metrics(
		id SERIAL PRIMARY KEY,
		node TEXT,
		cpu REAL,
		memory REAL,
		disk REAL,
		load REAL,
		network_in BIGINT,
		network_out BIGINT,
		timestamp BIGINT
	)
	`
	db.Exec(query)
}

func InsertMetric(db *sql.DB, m Metric) {
	query := `INSERT INTO metrics(node,cpu,memory,disk,load,network_in,network_out,timestamp) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	db.Exec(query, m.Node, m.CPU, m.Memory, m.Disk, m.Load, m.NetworkIn, m.NetworkOut, m.Timestamp)
}
