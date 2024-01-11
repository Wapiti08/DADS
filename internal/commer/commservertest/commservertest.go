package main

import (
	"distributed_anomaly_detection_system/internal/commer"
)

func main() {
	commer.Run(":2300")
}
