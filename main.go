package main

import (
	"github.com/amirhnajafiz/nba-bot/internal/bot"
	"github.com/amirhnajafiz/nba-bot/internal/metric"
)

func main() {
	metric.NewServer(metric.Config{}).Start()

	b, err := bot.New()
	if err != nil {
		panic(err)
	}

	b.Start()
}
