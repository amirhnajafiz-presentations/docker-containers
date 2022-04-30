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

func NewMetrics() Metrics {
	return Metrics{
		ConnectionErrors: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "connection_errors_total",
			Help:        "total number of connection errors",
			ConstLabels: nil,
		}),
		SuccessfulRequest: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "successful_request_total",
			Help:        "total number of successful requests",
			ConstLabels: nil,
		}),
		FailedRequest: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "failed_request_total",
			Help:        "total number of failed requests",
			ConstLabels: nil,
		}),
	}
}
