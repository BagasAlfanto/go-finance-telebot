package database

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	botToken string
	teleID   int64
)

func init() {
	botToken = os.Getenv("BOT_TOKEN")
	teleID, _ = strconv.ParseInt(os.Getenv("TELE_ID"), 10, 64)
}

func Connect() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	return bot
}

func CheckTelegramID(id int64) bool {
	return id == teleID
}
