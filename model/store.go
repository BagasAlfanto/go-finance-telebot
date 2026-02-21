package model

import "finance-telebot/database"

func StorePemasukan(p *Pemasukan) error {
	_, err := database.DB.Exec(
		"INSERT INTO pemasukan (nominal, deskripsi, kategori) VALUES (?, ?, ?)",
		p.Nominal, p.Deskripsi, p.Kategori,
	)
	return err
}

func StorePengeluaran(p *Pengeluaran) error {
	_, err := database.DB.Exec(
		"INSERT INTO pengeluaran (nominal, deskripsi, kategori) VALUES (?, ?, ?)",
		p.Nominal, p.Deskripsi, p.Kategori,
	)
	return err
}
