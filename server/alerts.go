package main

import (
	"fmt"
)

func CheckAlerts(m Metric) {
	if m.CPU > 90 {
		fmt.Println("ALERT CPU HIGH:", m.Node, m.CPU)
	}
	if m.Memory > 85 {
		fmt.Println("ALERT MEMORY HIGH:", m.Node, m.Memory)
	}
	if m.Disk > 90 {
		fmt.Println("ALERT DISK HIGH:", m.Node, m.Disk)
	}
}
