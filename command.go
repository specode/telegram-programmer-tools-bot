package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleCommand(id int64, msgId int, command string, args []string) tgbotapi.Chattable {
	log.Printf("[handleCommand] Command[%s] args: %v", command, args)

	var reply tgbotapi.Chattable
	switch command {
	case "start":
		reply = tgbotapi.NewMessage(id, MsgWelcome)
	case "help":
		reply = NewHelpMsg(id, msgId)
	case "md5":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}

		result := tgbotapi.NewMessage(id, md5enc(args[0]))
		result.ReplyToMessageID = msgId
		reply = result
	case "base64enc":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}
		encoded := base64.StdEncoding.EncodeToString([]byte(args[0]))
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "base64dec":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}
		decoded, _ := base64.StdEncoding.DecodeString(args[0])
		result := tgbotapi.NewMessage(id, string(decoded))
		result.ReplyToMessageID = msgId
		reply = result
	case "urlenc":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}

		encoded := url.QueryEscape(args[0])
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "urldec":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}

		decoded, _ := url.QueryUnescape(args[0])
		result := tgbotapi.NewMessage(id, decoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "htmlenc":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}

		encoded := html.EscapeString(args[0])
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "htmldec":
		if !checkArgs(args, 1) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}

		encoded := html.UnescapeString(args[0])
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "time2timestamp":
		var err error
		t := time.Now()
		if len(args) == 2 && len(args[0]) != 0 {
			t, err = time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", args[0], args[1]))
			if err != nil {
				reply = NewInvaildMsg(id, msgId, "time format wrong. Please input like 2016-01-02 19:00:01")
				break
			}
		}

		ts := t.Unix()

		result := tgbotapi.NewMessage(id, strconv.Itoa(int(ts)))
		result.ReplyToMessageID = msgId
		reply = result
	case "timestamp2time":
		if !checkArgs(args, 0) {
			reply = NewInvaildMsg(id, msgId, "please input arg.")
			break
		}

		ts, _ := strconv.ParseInt(args[0], 10, 64)
		t := time.Unix(ts, 0)

		result := tgbotapi.NewMessage(id, t.Format("2006-01-02 15:04:05"))
		result.ReplyToMessageID = msgId
		reply = result
	default:
		invaild := tgbotapi.NewMessage(id, ErrInvaild)
		invaild.ReplyToMessageID = msgId
		reply = invaild
	}

	return reply
}

func checkArgs(args []string, length int) bool {
	if len(args) != length || len(args[0]) == 0 {
		return false
	}

	return true
}

func md5enc(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
