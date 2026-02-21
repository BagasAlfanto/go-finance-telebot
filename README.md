# Finance Telebot

Bot Telegram untuk mencatat pemasukan dan pengeluaran keuangan pribadi, dengan penyimpanan menggunakan SQLite.

## Fitur

- Catat pemasukan dengan kategori
- Catat pengeluaran dengan kategori
- Rekap transaksi terakhir beserta total dan saldo
- Akses terbatas hanya untuk satu Telegram ID

## Perintah Bot

| Perintah | Contoh | Keterangan |
|---|---|---|
| `/masuk [nominal] [deskripsi] #[kategori]` | `/masuk 300000 uang pendaftaran #event` | Catat pemasukan |
| `/keluar [nominal] [deskripsi] #[kategori]` | `/keluar 50000 beli komponen #proyek` | Catat pengeluaran |
| `/rekap` | `/rekap` | Rekap 10 transaksi terakhir |
| `/rekap [jumlah]` | `/rekap 20` | Rekap N transaksi terakhir |
| `/rekap #[kategori]` | `/rekap #event` | Rekap berdasarkan kategori |
| `/rekap [jumlah] #[kategori]` | `/rekap 20 #event` | Kombinasi keduanya |

## Prasyarat

- [Go](https://go.dev/dl/) 1.21 atau lebih baru
- Akun Telegram dan bot token dari [@BotFather](https://t.me/BotFather)

## Cara Menjalankan

### 1. Clone repositori

```bash
git clone https://github.com/username/finance-telebot.git
cd finance-telebot
```

### 2. Install dependensi

```bash
go mod tidy
```

### 3. Copy .env.example ke .env

```bash
cp .env.example .env
```

### 4. Siapkan environment variable

Bot ini membutuhkan dua environment variable:

| Variable | Keterangan |
|---|---|
| `BOT_TOKEN` | Token bot dari @BotFather |
| `TELE_ID` | Telegram ID kamu (untuk membatasi akses) |


### 5. Jalankan bot

```bash
go run main.go
```

Atau build terlebih dahulu:

```bash
go build -o finance-telebot
./finance-telebot        # Linux/macOS
.\finance-telebot.exe    # Windows
```

## Catatan Keamanan

Bot ini hanya menerima pesan dari satu Telegram ID yang ditentukan via `TELE_ID`. Pesan dari ID lain akan langsung ditolak.
