package model

import (
	"finance-telebot/database"
	"time"
)

type Transaksi struct {
	Tipe      string 
	Nominal   int64
	Deskripsi string
	Kategori  string
	CreatedAt time.Time
}

func GetRekap(q *RekapQuery) ([]Transaksi, error) {
	var args []any

	filterKategori := ""
	if q.Kategori != "" {
		filterKategori = " AND kategori = ?"
	}

	query := `
		SELECT 'masuk' AS tipe, nominal, deskripsi, kategori, created_at
		FROM pemasukan
		WHERE 1=1` + filterKategori + `
		UNION ALL
		SELECT 'keluar' AS tipe, nominal, deskripsi, kategori, created_at
		FROM pengeluaran
		WHERE 1=1` + filterKategori + `
		ORDER BY created_at DESC
		LIMIT ?`

	if q.Kategori != "" {
		args = append(args, q.Kategori, q.Kategori)
	}
	args = append(args, q.Limit)

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []Transaksi
	for rows.Next() {
		var t Transaksi
		if err := rows.Scan(&t.Tipe, &t.Nominal, &t.Deskripsi, &t.Kategori, &t.CreatedAt); err != nil {
			return nil, err
		}
		hasil = append(hasil, t)
	}
	return hasil, rows.Err()
}
