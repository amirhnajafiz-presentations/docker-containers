package main

import (
	"github.com/amirhnajafiz/nba-bot/internal/bot"
	"github.com/amirhnajafiz/nba-bot/internal/config"
	"github.com/amirhnajafiz/nba-bot/internal/metric"
)

func main() {
	cfg := config.Load()

	metric.NewServer(cfg.Metric).Start()

	b, err := bot.New(cfg.Bot)
	if err != nil {
		panic(err)
	}

	b.Start()
}
