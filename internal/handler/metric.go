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

func newCounter(counterOpts prometheus.CounterOpts) prometheus.Counter {
	ev := prometheus.NewCounter(counterOpts)

	if err := prometheus.Register(ev); err != nil {
		panic(err)
	}

	return ev
}
