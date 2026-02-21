package controller

import (
	"finance-telebot/database"
	"finance-telebot/model"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StorePemasukan(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := update.Message
	if message == nil {
		return
	}

	text := message.Text
	if text == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Pesan tidak boleh kosong")
		database.SendAndLog(bot, msg)
		return
	}

	pemasukan, err := model.ParsePemasukan(text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Format pemasukan tidak valid: "+err.Error()+
			"\n\nFormat: <code>/masuk [nominal] [deskripsi] #[kategori]</code>"+
			"\nContoh: <code>/masuk 300000 joki tugas #job</code>")
		msg.ParseMode = "HTML"
		database.SendAndLog(bot, msg)
		return
	}

	err = model.StorePemasukan(pemasukan)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Gagal menyimpan pemasukan: "+err.Error())
		database.SendAndLog(bot, msg)
		return
	}

	responseText := fmt.Sprintf(
		"✅ Pemasukan sebesar <b>%s</b> berhasil disimpan!\n\n📝 Deskripsi: %s\n🏷 Kategori: #%s",
		model.FormatRupiah(pemasukan.Nominal),
		pemasukan.Deskripsi,
		pemasukan.Kategori,
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, responseText)
	msg.ReplyToMessageID = message.MessageID
	msg.ParseMode = "HTML"
	database.SendAndLog(bot, msg)
}
