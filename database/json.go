package database

import (
	"encoding/json"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const logFile = "./logs/chatlog.json"
const botLogFile = "./logs/botlog.json"

type ChatLog struct {
	Timestamp string `json:"timestamp"`
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	ChatID    int64  `json:"chat_id"`
	Command   string `json:"command,omitempty"`
	Message   string `json:"message"`
}

type BotLog struct {
	Timestamp string `json:"timestamp"`
	ChatID    int64  `json:"chat_id"`
	ReplyToID int    `json:"reply_to_id,omitempty"`
	Message   string `json:"message"`
}

func LogChat(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	entry := ChatLog{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		UserID:    update.Message.From.ID,
		Username:  update.Message.From.UserName,
		FirstName: update.Message.From.FirstName,
		ChatID:    update.Message.Chat.ID,
		Message:   update.Message.Text,
	}

	if len(entry.Message) > 0 && entry.Message[0] == '/' {
		end := len(entry.Message)
		for i, c := range entry.Message {
			if i > 0 && (c == ' ' || c == '\n') {
				end = i
				break
			}
		}
		entry.Command = entry.Message[:end]
	}

	_ = os.MkdirAll("./logs", 0755)

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	line, err := json.Marshal(entry)
	if err != nil {
		return
	}
	f.Write(line)
	f.WriteString("\n")
}

func SendAndLog(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	bot.Send(msg)

	entry := BotLog{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		ChatID:    msg.ChatID,
		ReplyToID: msg.ReplyToMessageID,
		Message:   msg.Text,
	}

	_ = os.MkdirAll("./logs", 0755)

	f, err := os.OpenFile(botLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	line, err := json.Marshal(entry)
	if err != nil {
		return
	}
	f.Write(line)
	f.WriteString("\n")
}
