package main

import "github.com/amirhnajafiz/nba-bot/internal/bot"

func main() {
	b, err := bot.New()
	if err != nil {
		panic(err)
	}

	b.Start()
}
