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

// TotalKategori menyimpan ringkasan total per kategori.
type TotalKategori struct {
	Kategori string
	Masuk    int64
	Keluar   int64
}

func (t TotalKategori) Saldo() int64 { return t.Masuk - t.Keluar }

// GetTotal mengembalikan total pemasukan & pengeluaran per kategori,
// diurutkan berdasarkan nama kategori.
func GetTotal(filterKategori string) ([]TotalKategori, error) {
	query := `
		SELECT kategori,
		       SUM(CASE WHEN tipe = 'masuk'  THEN nominal ELSE 0 END) AS masuk,
		       SUM(CASE WHEN tipe = 'keluar' THEN nominal ELSE 0 END) AS keluar
		FROM (
			SELECT 'masuk'  AS tipe, nominal, kategori FROM pemasukan
			UNION ALL
			SELECT 'keluar' AS tipe, nominal, kategori FROM pengeluaran
		)
		WHERE 1=1`

	var args []any
	if filterKategori != "" {
		query += " AND kategori = ?"
		args = append(args, filterKategori)
	}
	query += " GROUP BY kategori ORDER BY kategori"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []TotalKategori
	for rows.Next() {
		var t TotalKategori
		if err := rows.Scan(&t.Kategori, &t.Masuk, &t.Keluar); err != nil {
			return nil, err
		}
		hasil = append(hasil, t)
	}
	return hasil, rows.Err()
}
