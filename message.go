package main

import (
	"fmt"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	MsgWelcome    string   = "Welcome, This is programmer tools, /help"
	MsgHelpFormat string   = "Tools:\n\n%v"
	MsgCommand    []string = []string{
		"/md5 - Generate md5 encrypted string",
		"/base64enc - Base64 Encode",
		"/base64dec - Base64 Decode",
		"/urlenc - URL Encode",
		"/urldec - URL Decode",
		"/htmlenc - HTML Encode",
		"/htmldec - HTML Decode",
		"/time2timestamp - Convert time(YYYY-MM-DD HH:ii:ss) to timestamp",
		"/timestamp2time - Convert timestamp to time(YYYY-MM-DD HH:ii:ss)",
		"/help - Show help",
	}
)

func NewInvaildMsg(id int64, replyId int, desc string) tgbotapi.MessageConfig {
	invaild := tgbotapi.NewMessage(id, fmt.Sprintf("%s: %s", ErrInvaild, desc))
	invaild.ReplyToMessageID = replyId
	return invaild
}

func NewHelpMsg(id int64, replyId int) tgbotapi.MessageConfig {
	help := tgbotapi.NewMessage(id, fmt.Sprintf(MsgHelpFormat, strings.Join(MsgCommand, "\n")))
	help.ReplyToMessageID = replyId
	return help
}
