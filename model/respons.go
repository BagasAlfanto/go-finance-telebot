package model

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WelcomeMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "📊 <b>SELAMAT DATANG DI FINBOT!</b> 📊\n<i>Asisten Keuangan Pribadi Bagas</i> 💸\n━━━━━━━━━━━━━━━━━━━━━━\n\nHalo Bagas! 👋 Bot ini dikunci secara privat khusus untukmu. Mulai sekarang, mencatat pengeluaran proyek atau jajan harian cukup dengan sekali ketik.\n\n✨ <b>CARA PENGGUNAAN:</b>\n\n📤 <b>Catat Pengeluaran:</b>\n<code>/keluar [nominal] [deskripsi] #[kategori]</code>\nContoh: <code>/keluar 50000 beli komponen #nightguardian</code>\n\n📥 <b>Catat Pemasukan:</b>\n<code>/masuk [nominal] [deskripsi] #[kategori]</code>\nContoh: <code>/masuk 300000 uang pendaftaran #ceniptour</code>\n\n📈 <b>Tarik Laporan (Rekap):</b>\n<code>/rekap #[kategori]</code>\nContoh: <code>/rekap #insting</code>\n\nKetik langsung perintah di atas, dan database akan otomatis mencatatnya! 🚀"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
