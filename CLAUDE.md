# Project Overview: Sisprese (Mini Attendance API)

Sistem RESTful API presensi sederhana untuk latihan Golang & Gin Framework.

## Tech Stack
- Language: Go (Golang) v1.27+
- Framework: Gin Gonic (`github.com/gin-gonic/gin`)
- Storage: In-Memory (Struct / Slice Go, tanpa database external dulu)

## Data Structure

### Employee (Karyawan)
- ID (string, e.g., "EMP001")
- Name (string)
- Department (string)

### Attendance (Presensi)
- ID (string)
- EmployeeID (string)
- Type (string: "IN" / "OUT")
- Timestamp (time.Time)

## Endpoints Required
1. `GET /health` -> Cek status server.
2. `POST /employees` -> Tambah karyawan baru.
3. `GET /employees` -> Ambil daftar semua karyawan.
4. `POST /attendance/clock-in` -> Catat presensi masuk (Body: `employee_id`).
5. `POST /attendance/clock-out` -> Catat presensi keluar (Body: `employee_id`).
6. `GET /attendance` -> Ambil semua riwayat presensi.

## Project Rules
- Gunakan struktur folder yang rapi (handler, model, router).
- Tambahkan middleware logger sederhana untuk mencatat setiap request HTTP yang masuk.
- Sertakan penanganan error yang baik (response JSON dengan status code HTTP yang sesuai).