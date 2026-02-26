package controller

import (
	"finance-telebot/database"
	"finance-telebot/model"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Total menampilkan ringkasan total pemasukan & pengeluaran per kategori.
// Penggunaan:
//
//	/total          → semua kategori
//	/total #nama    → kategori tertentu
func Total(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := update.Message
	if message == nil {
		return
	}

	// Parsing opsional filter kategori dari teks perintah
	filterKategori := ""
	text := strings.TrimSpace(strings.TrimPrefix(message.Text, "/total"))
	if idx := strings.LastIndex(text, "#"); idx != -1 {
		filterKategori = strings.TrimSpace(text[idx+1:])
	}

	data, err := model.GetTotal(filterKategori)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Gagal mengambil data: "+err.Error())
		database.SendAndLog(bot, msg)
		return
	}

	if len(data) == 0 {
		teks := "📭 Belum ada transaksi"
		if filterKategori != "" {
			teks += fmt.Sprintf(" untuk kategori <b>#%s</b>", filterKategori)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, teks)
		msg.ParseMode = "HTML"
		database.SendAndLog(bot, msg)
		return
	}

	var sb strings.Builder
	judul := "📊 <b>Rekap Total Keuangan</b>"
	if filterKategori != "" {
		judul += fmt.Sprintf(" — <b>#%s</b>", filterKategori)
	}
	sb.WriteString(judul + "\n")
	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━\n")

	var grandMasuk, grandKeluar int64
	for _, d := range data {
		saldo := d.Saldo()
		saldoStr := model.FormatRupiah(saldo)
		saldoIcon := "💰"
		if saldo < 0 {
			saldoIcon = "🔴"
		}

		sb.WriteString(fmt.Sprintf(
			"\n🏷 <b>#%s</b>\n  📥 Masuk : %s\n  📤 Keluar: %s\n  %s Saldo : <b>%s</b>\n",
			d.Kategori,
			model.FormatRupiah(d.Masuk),
			model.FormatRupiah(d.Keluar),
			saldoIcon,
			saldoStr,
		))

		grandMasuk += d.Masuk
		grandKeluar += d.Keluar
	}

	grandSaldo := grandMasuk - grandKeluar
	grandIcon := "💰"
	if grandSaldo < 0 {
		grandIcon = "🔴"
	}

	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━\n")
	sb.WriteString(fmt.Sprintf("📥 <b>Grand Total Masuk : %s</b>\n", model.FormatRupiah(grandMasuk)))
	sb.WriteString(fmt.Sprintf("📤 <b>Grand Total Keluar: %s</b>\n", model.FormatRupiah(grandKeluar)))
	sb.WriteString(fmt.Sprintf("%s <b>Grand Saldo       : %s</b>", grandIcon, model.FormatRupiah(grandSaldo)))

	msg := tgbotapi.NewMessage(message.Chat.ID, sb.String())
	msg.ReplyToMessageID = message.MessageID
	msg.ParseMode = "HTML"
	database.SendAndLog(bot, msg)
}
