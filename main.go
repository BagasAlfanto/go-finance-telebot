package main

import (
	"finance-telebot/controller"
	"finance-telebot/database"
	"finance-telebot/model"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	database.InitDB()

	bot := database.Connect()
	if bot == nil {
		log.Panic("Failed to connect to the bot API")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		senderID := update.Message.From.ID
		if !database.CheckTelegramID(senderID) {
			log.Printf("Unauthorized access attempt from ID: %d", senderID)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Anda tidak memiliki akses")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			continue
		}

		text := update.Message.Text
		switch {
		case text == "/start":
			model.WelcomeMessage(bot, update)
		case strings.HasPrefix(text, "/masuk"):
			controller.StorePemasukan(bot, update)
		case strings.HasPrefix(text, "/keluar"):
			controller.StorePengeluaran(bot, update)
		case strings.HasPrefix(text, "/rekap"):
			controller.Recap(bot, update)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Perintah tidak dikenali. Ketik /help untuk bantuan.")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, text)
	}
}
