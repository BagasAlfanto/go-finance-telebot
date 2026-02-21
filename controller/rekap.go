package controller

import (
	"finance-telebot/database"
	"finance-telebot/model"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Recap(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := update.Message
	if message == nil {
		return
	}

	q, err := model.ParseRekap(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Format tidak valid: "+err.Error()+
			"\n\nFormat: <code>/rekap [jumlah] #[kategori]</code>"+
			"\nContoh:"+
			"\n  <code>/rekap</code>"+
			"\n  <code>/rekap 20</code>"+
			"\n  <code>/rekap #nightguardian</code>"+
			"\n  <code>/rekap 20 #nightguardian</code>")
		msg.ParseMode = "HTML"
		database.SendAndLog(bot, msg)
		return
	}

	transaksi, err := model.GetRekap(q)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Gagal mengambil data: "+err.Error())
		database.SendAndLog(bot, msg)
		return
	}

	if len(transaksi) == 0 {
		teks := "📭 Belum ada transaksi"
		if q.Kategori != "" {
			teks += fmt.Sprintf(" untuk kategori <b>#%s</b>", q.Kategori)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, teks)
		msg.ParseMode = "HTML"
		database.SendAndLog(bot, msg)
		return
	}

	judul := fmt.Sprintf("📊 <b>%d Transaksi Terakhir</b>", q.Limit)
	if q.Kategori != "" {
		judul += fmt.Sprintf(" — <b>#%s</b>", q.Kategori)
	}

	var totalMasuk, totalKeluar int64
	var sb strings.Builder
	sb.WriteString(judul + "\n")
	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━\n")

	for i, t := range transaksi {
		var ikon, tipeLabel string
		if t.Tipe == "masuk" {
			ikon = "📥"
			tipeLabel = "Masuk"
			totalMasuk += t.Nominal
		} else {
			ikon = "📤"
			tipeLabel = "Keluar"
			totalKeluar += t.Nominal
		}

		sb.WriteString(fmt.Sprintf(
			"\n%d. %s <b>%s</b> — %s\n    📝 %s  🏷 #%s\n    🕐 %s\n",
			i+1,
			ikon,
			tipeLabel,
			model.FormatRupiah(t.Nominal),
			t.Deskripsi,
			t.Kategori,
			t.CreatedAt.Format("02 Jan 2006, 15:04"),
		))
	}

	sb.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━\n")
	sb.WriteString(fmt.Sprintf("📥 Total Masuk : <b>%s</b>\n", model.FormatRupiah(totalMasuk)))
	sb.WriteString(fmt.Sprintf("📤 Total Keluar: <b>%s</b>\n", model.FormatRupiah(totalKeluar)))
	saldo := totalMasuk - totalKeluar
	saldoLabel := "💰 Saldo"
	if saldo < 0 {
		saldoLabel = "🔴 Saldo"
	}
	sb.WriteString(fmt.Sprintf("%s       : <b>%s</b>", saldoLabel, model.FormatRupiah(saldo)))

	msg := tgbotapi.NewMessage(message.Chat.ID, sb.String())
	msg.ReplyToMessageID = message.MessageID
	msg.ParseMode = "HTML"
	database.SendAndLog(bot, msg)
}
