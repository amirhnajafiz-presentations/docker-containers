package handler

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "tel-bot"
	Subsystem = "cli-srv"
)

// Metrics has all the client metrics.
type Metrics struct {
	ConnectionErrors  prometheus.Counter
	SuccessfulRequest prometheus.Counter
	FailedRequest     prometheus.Counter
}
