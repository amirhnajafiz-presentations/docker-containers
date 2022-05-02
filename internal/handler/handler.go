package handler

import (
	"log"

	"github.com/amirhnajafiz/nba-bot/internal/api"
	"gopkg.in/telebot.v3"
)

type Handler struct {
	Metric Metrics
}

func (h Handler) Test(context telebot.Context) error {
	h.Metric.SuccessfulRequest.Add(1)

	return context.Send("Success")
}

func (h Handler) Request(context telebot.Context) error {
	h.Metric.SuccessfulRequest.Add(1)

	r, s := api.MakeRequest()

	log.Println(r)

	return context.Send(s)
}
