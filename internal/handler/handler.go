package handler

import (
	"gopkg.in/telebot.v3"
)

type Handler struct {
	Metric Metrics
}

func (h Handler) Test(context telebot.Context) error {
	h.Metric.SuccessfulRequest.Add(1)

	return context.Send("Success")
}
