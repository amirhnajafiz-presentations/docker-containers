package bot

import "gopkg.in/telebot.v3"

type Bot struct {
}

func New() (*telebot.Bot, error) {
	b, err := telebot.NewBot(Config())
	if err != nil {
		return nil, err
	}

	// TODO: register handler

	return b, nil
}
