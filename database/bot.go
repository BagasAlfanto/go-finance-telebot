package database

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	botToken string
	teleID   int64
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[info] .env file tidak ditemukan, menggunakan environment variable sistem")
	}
	botToken = os.Getenv("BOT_TOKEN")
	teleID, _ = strconv.ParseInt(os.Getenv("TELE_ID"), 10, 64)
	if botToken == "" {
		log.Panic("BOT_TOKEN tidak valid. Isi file .env atau set environment variable BOT_TOKEN")
	}
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
