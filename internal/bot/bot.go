package bot

import (
	"github.com/amirhnajafiz/nba-bot/internal/handler"
	"gopkg.in/telebot.v3"
)

func New(config Config) (*telebot.Bot, error) {
	// load bot configs
	b, err := telebot.NewBot(LoadConfigs(config))
	if err != nil {
		return nil, err
	}

	// create the handler
	h := handler.Handler{
		Metric: handler.NewMetrics(),
	}

	b.Handle("/test", h.Test)
	b.Handle("/view", h.Request)

	return b, nil
}
