package model

import "finance-telebot/database"

func GetKategori() ([]string, error) {
	rows, err := database.DB.Query("SELECT DISTINCT kategori FROM (SELECT kategori FROM pemasukan UNION ALL SELECT kategori FROM pengeluaran) ORDER BY kategori")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []string
	for rows.Next() {
		var k string
		if err := rows.Scan(&k); err != nil {
			return nil, err
		}
		list = append(list, k)
	}
	return list, rows.Err()
}
