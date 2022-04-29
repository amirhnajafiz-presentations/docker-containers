package bot

import (
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func Config() telebot.Settings {
	return telebot.Settings{
		Token: os.Getenv("token"),
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}
}
