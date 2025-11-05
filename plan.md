# Rencana Pembuatan Tool Sederhana + Linter & LSP

## Tujuan
- Menyediakan bukti hasil pembelajaran berupa sebuah tool CLI sederhana berbasis Go yang dapat dieksekusi dengan cepat.
- Mengintegrasikan linter dan LSP sesuai best practice agar proyek siap dikembangkan lebih lanjut.

## Ruang Lingkup
- Bahasa: Go (selaras dengan struktur repo `Gos/` dan dukungan ekosistem tooling).
- Fungsi tool: utilitas CLI ringan (mis. membaca file konfigurasi NiFi dan menampilkan ringkasan sederhana) untuk demonstrasi.
- Linter: `golangci-lint` dengan konfigurasi minimal tetapi mengikuti rekomendasi komunitas.
- LSP: `gopls` dengan pengaturan dasar agar editor dapat langsung menggunakan informasi proyek.

## Deliverable Utama
1. Struktur proyek Go modular (`cmd/`, `internal/` atau `pkg/`) beserta `go.mod`.
2. Implementasi fitur CLI minimal yang menegaskan alur input → pemrosesan → output.
3. Konfigurasi `golangci-lint.yaml` dan integrasi ke workflow (manual command + hook opsional).
4. Pengaturan `gopls` via `go.work`/`settings.json` contoh, atau penjelasan cara aktivasi jika editor spesifik.
5. Dokumentasi singkat (README) mengenai cara menjalankan tool, linter, dan LSP.

## Pendekatan Pengerjaan (Tingkat Tinggi)
1. **Bootstrap Proyek**
   - Inisialisasi `go mod init`.
   - Siapkan struktur direktori standar (`cmd/toolname/main.go`, `internal/app`).
2. **Implementasi Fitur Minimal**
   - Buat paket yang membaca file (dummy JSON/CSV) dan menampilkan metrik.
   - Tambahkan flag CLI (`cobra` opsional; jika waktu sempit gunakan `flag`).
3. **Quality Tooling**
   - Tambahkan `golangci-lint` (local config + panduan instalasi cepat).
   - Konfigurasikan target linter (mis. `govet`, `staticcheck`, `errcheck`).
4. **LSP Enablement**
   - Pastikan `gopls` terdeteksi (`go env GOPATH` + modul).
   - Cantumkan petunjuk integrasi di README (VS Code/Neovim minimal).
5. **Testing & Contoh Output**
   - Tambahkan unit test ringkas untuk fungsi utama.
   - Sertakan contoh perintah `go test`, `go run`, `golangci-lint run`.
6. **Dokumentasi**
   - Perbarui README dengan langkah setup, jalankan, lint, debugging.

## Best Practice yang Harus Diikuti
- Pertahankan struktur modul Go idiomatik.
- Pisahkan logic bisnis dari kode CLI (paket internal).
- Gunakan dependency minimal; hindari menambah library berat bila tidak perlu.
- Coverage test dasar untuk fungsi utama (>1 unit test).
- Pastikan lint & test dijalankan sebelum menyatakan selesai.

## Jadwal Singkat (estimasi cepat)
| Tahap | Estimasi | Catatan |
| --- | --- | --- |
| Bootstrap proyek | 30 mnt | Inisialisasi modul & struktur folder |
| Implementasi fitur | 60 mnt | Termasuk unit test dasar |
| Setup lint & LSP | 45 mnt | Konfigurasi `golangci-lint`, verifikasi `gopls` |
| Dokumentasi akhir | 30 mnt | Update README + contoh perintah |
| Buffer | 15 mnt | Menangani bug kecil / lint issue |

Total target: ± 3 jam pengerjaan fokus.

## Definition of Done (DoD)
- Proyek Go dapat di-`go run` dan menghasilkan output sesuai spesifikasi sederhana.
- `go test ./...` dan `golangci-lint run` lulus tanpa error.
- Petunjuk aktivasi `gopls` tersedia (README mencakup langkah editor umum).
- Dokumentasi ringkas menjelaskan tujuan tool, cara pakai, dan alur kontribusi dasar.
- Semua file konfigurasi (lint, modul) tersusun rapi di repo.

## Langkah Berikutnya Setelah Plan Disetujui
1. Eksekusi bootstrap + commit awal.
2. Iterasi implementasi sesuai tahapan di atas.
3. Validasi terhadap DoD dan lakukan penyesuaian bila ada masukan.
