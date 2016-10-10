package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func NewBot(token string, debug bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = debug

	return &Bot{API: bot}, nil
}

func (bot *Bot) Serve(routeFunc func(data tgbotapi.Update)) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.API.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go routeFunc(update)
	}

	return nil
}
