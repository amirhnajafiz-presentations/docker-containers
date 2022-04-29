package handler

import (
	"gopkg.in/telebot.v3"
)

type Handler struct {
}

func (h Handler) Test(context telebot.Context) error {
	// TODO: create a test route
	return context.Send("Hello")
}
