package bot

import (
	"github.com/amirhnajafiz/nba-bot/internal/handler"
	"gopkg.in/telebot.v3"
)

type Bot struct {
}

func New() (*telebot.Bot, error) {
	b, err := telebot.NewBot(LoadConfigs(Config{}))
	if err != nil {
		return nil, err
	}

	// TODO: register handler
	h := handler.Handler{}

	b.Handle("/hello", h.Test)

	return b, nil
}
