package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Pemasukan struct {
	Nominal   int64
	Deskripsi string
	Kategori  string
}

func ParsePemasukan(text string) (*Pemasukan, error) {
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "/masuk")
	text = strings.TrimSpace(text)

	if text == "" {
		return nil, errors.New("format: /masuk [nominal] [deskripsi] #[kategori]")
	}

	hashIdx := strings.LastIndex(text, "#")
	if hashIdx == -1 {
		return nil, errors.New("kategori tidak ditemukan, gunakan #kategori di akhir")
	}

	kategori := strings.TrimSpace(text[hashIdx+1:])
	if kategori == "" {
		return nil, errors.New("kategori tidak boleh kosong")
	}

	text = strings.TrimSpace(text[:hashIdx])

	parts := strings.SplitN(text, " ", 2)
	if len(parts) < 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
		return nil, fmt.Errorf("format: /masuk [nominal] [deskripsi] #[kategori]")
	}

	nominalStr := strings.ReplaceAll(parts[0], ".", "")
	nominal, err := strconv.ParseInt(nominalStr, 10, 64)
	if err != nil || nominal <= 0 {
		return nil, fmt.Errorf("nominal tidak valid: %s", parts[0])
	}

	return &Pemasukan{
		Nominal:   nominal,
		Deskripsi: strings.TrimSpace(parts[1]),
		Kategori:  kategori,
	}, nil
}

type Pengeluaran struct {
	Nominal   int64
	Deskripsi string
	Kategori  string
}

func ParsePengeluaran(text string) (*Pengeluaran, error) {
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "/keluar")
	text = strings.TrimSpace(text)

	if text == "" {
		return nil, errors.New("format: /keluar [nominal] [deskripsi] #[kategori]")
	}

	hashIdx := strings.LastIndex(text, "#")
	if hashIdx == -1 {
		return nil, errors.New("kategori tidak ditemukan, gunakan #kategori di akhir")
	}

	kategori := strings.TrimSpace(text[hashIdx+1:])
	if kategori == "" {
		return nil, errors.New("kategori tidak boleh kosong")
	}

	text = strings.TrimSpace(text[:hashIdx])

	parts := strings.SplitN(text, " ", 2)
	if len(parts) < 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
		return nil, fmt.Errorf("format: /keluar [nominal] [deskripsi] #[kategori]")
	}

	nominalStr := strings.ReplaceAll(parts[0], ".", "")
	nominal, err := strconv.ParseInt(nominalStr, 10, 64)
	if err != nil || nominal <= 0 {
		return nil, fmt.Errorf("nominal tidak valid: %s", parts[0])
	}

	return &Pengeluaran{
		Nominal:   nominal,
		Deskripsi: strings.TrimSpace(parts[1]),
		Kategori:  kategori,
	}, nil
}

type RekapQuery struct {
	Limit    int
	Kategori string 
}

func ParseRekap(text string) (*RekapQuery, error) {
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "/rekap")
	text = strings.TrimSpace(text)

	q := &RekapQuery{Limit: 10}

	if text == "" {
		return q, nil
	}

	if hashIdx := strings.LastIndex(text, "#"); hashIdx != -1 {
		q.Kategori = strings.TrimSpace(text[hashIdx+1:])
		text = strings.TrimSpace(text[:hashIdx])
	}

	if text != "" {
		limit, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil || limit <= 0 {
			return nil, fmt.Errorf("jumlah tidak valid: %s", text)
		}
		if limit > 50 {
			limit = 50
		}
		q.Limit = limit
	}

	return q, nil
}

func FormatRupiah(nominal int64) string {
	s := strconv.FormatInt(nominal, 10)
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return "Rp " + result
}
