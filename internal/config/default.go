package config

import "github.com/amirhnajafiz/nba-bot/internal/bot"

func Default() Config {
	return Config{
		Bot: bot.Config{
			Token: "",
		},
	}
}
