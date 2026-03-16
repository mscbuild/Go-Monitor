package main

type Metric struct {
	Node       string
	CPU        float64
	Memory     float64
	Disk       float64
	Load       float64
	NetworkIn  uint64
	NetworkOut uint64
	Timestamp  int64
}
