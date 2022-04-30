package bot

import (
	"github.com/amirhnajafiz/nba-bot/internal/handler"
	"gopkg.in/telebot.v3"
)

type Bot struct {
}

func New(config Config) (*telebot.Bot, error) {
	b, err := telebot.NewBot(LoadConfigs(config))
	if err != nil {
		return nil, err
	}

	h := handler.Handler{
		Metric: handler.NewMetrics(),
	}

	b.Handle("/hello", h.Test)

	return b, nil
}
