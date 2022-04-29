package bot

import (
	"time"

	"gopkg.in/telebot.v3"
)

type Config struct {
	Token string `koanf:"token"`
}

func LoadConfigs(cfg Config) telebot.Settings {
	return telebot.Settings{
		Token: cfg.Token,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}
}
