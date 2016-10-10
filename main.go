package main

import (
	"flag"
	"log"
)

var bot *Bot

func main() {
	var file string
	flag.StringVar(&file, "c", "config.yaml", "config file")
	flag.Parse()

	if err := initConfig(file); err != nil {
		log.Fatal(err)
	}

	var err error
	bot, err = NewBot(cfg.Bot.Token, cfg.Bot.Debug)
	if err != nil {
		panic(err)
	}

	log.Println("Bot Auth Success:", bot.API.Self.UserName)
	bot.Serve(routePackage)
}
