package main

import (
	"fmt"

	"github.com/amirhnajafiz/nba-bot/internal/bot"
	"github.com/amirhnajafiz/nba-bot/internal/config"
	"github.com/amirhnajafiz/nba-bot/internal/metric"
)

func main() {
	// load configs
	cfg := config.Load()

	// starting prometheus server
	metric.NewServer(cfg.Metric).Start()

	// creating new bot
	b, err := bot.New(cfg.Bot)
	if err != nil {
		panic(err)
	}

	// starting our bot
	fmt.Println("Waiting for telegram bot to start ...")
	b.Start()
}
