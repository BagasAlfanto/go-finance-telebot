package controller

import (
	"finance-telebot/database"
	"finance-telebot/model"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ListKategori(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := update.Message
	if message == nil {
		return
	}

	kategori, err := model.GetKategori()
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Gagal mengambil kategori: "+err.Error())
		database.SendAndLog(bot, msg)
		return
	}

	var text string
	if len(kategori) == 0 {
		text = "📭 Belum ada kategori tersimpan."
	} else {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("🏷 <b>%d Kategori Tersedia:</b>\n", len(kategori)))
		sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━\n")
		for i, k := range kategori {
			sb.WriteString(fmt.Sprintf("%d. <code>#%s</code>\n", i+1, k))
		}
		text = sb.String()
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	msg.ParseMode = "HTML"
	database.SendAndLog(bot, msg)
}
