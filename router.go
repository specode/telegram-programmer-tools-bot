package main

import (
	// "fmt"
	// "strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func routePackage(update tgbotapi.Update) {
	if (!update.Message.Chat.IsPrivate()) &&
		((update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup()) && !bot.API.IsMessageToMe(*update.Message)) {
		return
	}

	switch {
	case update.Message.IsCommand():
		reply := handleCommand(
			update.Message.Chat.ID,
			update.Message.MessageID,
			update.Message.Command(),
			update.Message.CommandArguments(),
		)

		bot.API.Send(reply)
	case update.Message.Text != "":
		bot.API.Send(NewHelpMsg(update.Message.Chat.ID, 0))
	default:
	}
}
