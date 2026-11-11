# Refactoring & Feature Expansion Plan: Sisprese Backend (Golang)

## Project Context

`sisprese` adalah aplikasi REST API presensi berbasis Golang dan Gin Framework. Saat ini kode masih menyimpan data di memori (in-memory slice) dan belum memiliki autentikasi serta persistensi database.

## Objective

Tingkatkan arsitektur proyek `sisprese` agar siap produksi dan lebih profesional untuk portofolio magang.

## Task Requirements for Codex

### 1. Database Integration (SQLite / GORM)

- Integrasikan ORM `gorm.io/gorm` dan driver SQLite (`gorm.io/driver/sqlite`).
- Buat file `config/database.go` untuk inisialisasi koneksi database dan auto-migrate model.
- Pindahkan penyimpanan data `Employee` dan `Attendance` dari slice memori ke tabel SQLite (`sisprese.db`).

### 2. Architecture & Directory Refactoring

Tata ulang struktur folder mengikuti standar Clean Architecture:

- `config/`: Inisialisasi database & environment variables.
- `models/`: Entity structs & database schema.
- `repositories/`: Akses data langsung ke database (GORM).
- `services/`: Business logic (validasi clock-in/clock-out, shift kerja).
- `handlers/`: HTTP Controllers (Gin binding & JSON response).
- `routes/`: Registrasi endpoint HTTP.

### 3. Additional Features

- **Filter Presensi:** Tambahkan query parameter pada `GET /attendance` untuk filter berdasarkan tanggal dan ID Karyawan.
- **Clock-In Note:** Tambahkan kolom/field `Note` (keterangan presensi/lokasi) pada struct `Attendance`.

## Instructions for Codex Execution

1. Analisis kode yang ada di `main.go`, `handlers/handler.go`, dan `models/model.go`.
2. Install dependency yang dibutuhkan via `go get` (GORM & SQLite).
3. Buat file dan struktur folder baru secara bertahap tanpa merusak fitur Clock-In/Clock-Out yang sudah berjalan.
4. Pastikan `go run main.go` dapat berjalan dengan sukses setelah refactoring.
