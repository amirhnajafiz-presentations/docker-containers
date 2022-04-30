package config

import (
	"github.com/amirhnajafiz/nba-bot/internal/bot"
	"github.com/amirhnajafiz/nba-bot/internal/metric"
)

func Default() Config {
	return Config{
		Bot: bot.Config{
			Token: "",
		},
		Metric: metric.Config{
			Host: "",
		},
	}
}
