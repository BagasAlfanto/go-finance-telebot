package controller

import (
	"finance-telebot/model"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StorePengeluaran(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := update.Message
	if message == nil {
		return
	}

	text := message.Text
	if text == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Pesan tidak boleh kosong")
		bot.Send(msg)
		return
	}

	pengeluaran, err := model.ParsePengeluaran(text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Format pengeluaran tidak valid: "+err.Error()+
			"\n\nFormat: <code>/keluar [nominal] [deskripsi] #[kategori]</code>"+
			"\nContoh: <code>/keluar 50000 ganti oli #service</code>")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		return
	}

	err = model.StorePengeluaran(pengeluaran)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Gagal menyimpan pengeluaran: "+err.Error())
		bot.Send(msg)
		return
	}

	responseText := fmt.Sprintf(
		"✅ Pengeluaran sebesar <b>%s</b> berhasil disimpan!\n\n📝 Deskripsi: %s\n🏷 Kategori: #%s",
		model.FormatRupiah(pengeluaran.Nominal),
		pengeluaran.Deskripsi,
		pengeluaran.Kategori,
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, responseText)
	msg.ReplyToMessageID = message.MessageID
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
