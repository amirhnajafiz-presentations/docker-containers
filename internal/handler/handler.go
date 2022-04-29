package handler

import (
	"gopkg.in/telebot.v3"
)

type Handler struct {
}

func (h Handler) Test(context telebot.Context) error {
	return context.Send("Success")
}
