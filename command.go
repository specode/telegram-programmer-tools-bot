package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	// "fmt"
	"html"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	ErrInvaild           = "Invaild"
	ErrEmptyParams       = "please input arg. eg: /md5 xxx"
	ErrInvaildTimeParams = "time format wrong. Please input like 2016-01-02 19:00:01"
)

func handleCommand(id int64, msgId int, command string, args string) tgbotapi.Chattable {
	log.Printf("[handleCommand] Command[%s] args: %v", command, args)

	var reply tgbotapi.Chattable
	switch command {
	case "start":
		reply = tgbotapi.NewMessage(id, MsgWelcome)
	case "help":
		reply = NewHelpMsg(id, msgId)
	case "md5":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}

		result := tgbotapi.NewMessage(id, md5enc(args))
		result.ReplyToMessageID = msgId
		reply = result
	case "base64enc":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}
		encoded := base64.StdEncoding.EncodeToString([]byte(args))
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "base64dec":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}
		decoded, _ := base64.StdEncoding.DecodeString(args)
		result := tgbotapi.NewMessage(id, string(decoded))
		result.ReplyToMessageID = msgId
		reply = result
	case "urlenc":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}

		encoded := url.QueryEscape(args)
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "urldec":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}

		decoded, _ := url.QueryUnescape(args)
		result := tgbotapi.NewMessage(id, decoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "htmlenc":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}

		encoded := html.EscapeString(args)
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "htmldec":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}

		encoded := html.UnescapeString(args)
		result := tgbotapi.NewMessage(id, encoded)
		result.ReplyToMessageID = msgId
		reply = result
	case "time2timestamp":
		var err error
		t := time.Now()
		if len(args) != 0 {
			t, err = time.Parse("2006-01-02 15:04:05", args)
			if err != nil {
				reply = NewInvaildMsg(id, msgId, ErrInvaildTimeParams)
				break
			}
		}

		ts := t.Unix()

		result := tgbotapi.NewMessage(id, strconv.Itoa(int(ts)))
		result.ReplyToMessageID = msgId
		reply = result
	case "timestamp2time":
		if len(args) == 0 {
			reply = NewInvaildMsg(id, msgId, ErrEmptyParams)
			break
		}

		ts, _ := strconv.ParseInt(args, 10, 64)
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
