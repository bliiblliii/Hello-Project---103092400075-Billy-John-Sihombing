package main

import (
	"fmt"
)

type Akun struct {
	ID     string
	Nama   string
	Saldo  float64
	Status string
}

type Transaksi struct {
	ID         string
	Jenis      string
	Jumlah     float64
	Keterangan string
	AkunID     string
}

const maxAkun = 100

var daftarAkun [maxAkun]Akun
var jumlahAkun int = 0

const maxTransaksi = 100

var daftarTransaksi [maxTransaksi]Transaksi
var jumlahTransaksi int = 0

var counterID int = 0

func buatID() string {
	counterID++
	return fmt.Sprintf("%d", counterID)
}

func cariAkun(id string) int { //sequential search
	for i := 0; i < jumlahAkun; i++ {
		if daftarAkun[i].ID == id {
			return i
		}
	}
	return -1
}

func cariAkunBinary(id string) int { //binary search
	left := 0
	right := jumlahAkun - 1

	for left <= right {
		mid := left + (right-left)/2

		if daftarAkun[mid].ID == id {
			return mid
		} else if daftarAkun[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func urutkanAkun(isAscending bool) { //insertion sort
	if jumlahAkun <= 1 {
		return
	}

	for i := 1; i < jumlahAkun; i++ {
		key := daftarAkun[i]
		j := i - 1

		for j >= 0 {
			if isAscending {
				if daftarAkun[j].Nama > key.Nama {
					daftarAkun[j+1] = daftarAkun[j]
					j = j - 1
				} else {
					return
				}
			} else {
				if daftarAkun[j].Nama < key.Nama {
					daftarAkun[j+1] = daftarAkun[j]
					j = j - 1
				} else {
					return
				}
			}
		}
		daftarAkun[j+1] = key
	}

	if isAscending {
		fmt.Println("Daftar akun berhasil diurutkan secara Ascending (berdasarkan Nama) menggunakan Insertion Sort.")
	} else {
		fmt.Println("Daftar akun berhasil diurutkan secara Descending (berdasarkan Nama) menggunakan Insertion Sort.")
	}
}

func registrasiAkun() {
	if jumlahAkun >= maxAkun {
		fmt.Println("Maaf, kapasitas akun penuh.")
		return
	}

	var akun Akun
	akun.ID = buatID()
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&akun.Nama)
	akun.Saldo = 0
	akun.Status = "Menunggu"

	daftarAkun[jumlahAkun] = akun
	jumlahAkun++

	fmt.Println("Registrasi berhasil. Menunggu persetujuan admin.")
}

func persetujuanAkun() {
	var id string
	fmt.Print("Masukkan ID Akun yang akan disetujui/ditolak: ")
	fmt.Scanln(&id)

	index := cariAkun(id)
	if index == -1 {
		fmt.Println("Akun tidak ditemukan.")
		return
	}

	fmt.Printf("Nama Akun: %s\n", daftarAkun[index].Nama)
	fmt.Printf("Status saat ini: %s\n", daftarAkun[index].Status)

	var pilihan int
	fmt.Print("1. Setujui\n2. Tolak\nPilih: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		daftarAkun[index].Status = "Aktif"
		fmt.Println("Akun disetujui.")
	} else if pilihan == 2 {
		daftarAkun[index].Status = "Ditolak"
		fmt.Println("Akun ditolak.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func tampilkanDaftarAkun() { // selection sort
	if jumlahAkun == 0 {
		fmt.Println("Belum ada akun terdaftar.")
		return
	}

	var tempAkun [maxAkun]Akun
	for i := 0; i < jumlahAkun; i++ {
		tempAkun[i] = daftarAkun[i]
	}

	for i := 0; i < jumlahAkun-1; i++ {
		min_idx := i
		for j := i + 1; j < jumlahAkun; j++ {
			if tempAkun[j].Nama < tempAkun[min_idx].Nama {
				min_idx = j
			}
		}
		tempAkun[i], tempAkun[min_idx] = tempAkun[min_idx], tempAkun[i]
	}

	fmt.Println("Daftar Akun (Diurutkan berdasarkan Nama menggunakan Selection Sort):")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%-15s%-20s%-10s%-20s\n", "ID", "Nama", "Saldo", "Status")
	fmt.Println("--------------------------------------------------")
	for i := 0; i < jumlahAkun; i++ {
		fmt.Printf("%-15s%-20s%-10.2f%-20s\n", tempAkun[i].ID, tempAkun[i].Nama, tempAkun[i].Saldo, tempAkun[i].Status)
	}
	fmt.Println("--------------------------------------------------")
}

func kirimUang() {
	var idPengirim string
	var idPenerima string
	var jumlah float64

	fmt.Print("Masukkan ID Akun Pengirim: ")
	fmt.Scanln(&idPengirim)
	fmt.Print("Masukkan ID Akun Penerima: ")
	fmt.Scanln(&idPenerima)
	fmt.Print("Masukkan Jumlah Uang yang Dikirim: ")
	fmt.Scanln(&jumlah)

	indexPengirim := cariAkun(idPengirim)
	indexPenerima := cariAkun(idPenerima)

	if indexPengirim == -1 || indexPenerima == -1 {
		fmt.Println("Akun pengirim atau penerima tidak ditemukan.")
		return
	}

	if daftarAkun[indexPengirim].Status != "Aktif" || daftarAkun[indexPenerima].Status != "Aktif" {
		fmt.Println("Pastikan kedua akun aktif untuk melakukan transfer.")
		return
	}

	if daftarAkun[indexPengirim].Saldo < jumlah {
		fmt.Println("Saldo tidak mencukupi.")
		return
	}

	daftarAkun[indexPengirim].Saldo -= jumlah
	daftarAkun[indexPenerima].Saldo += jumlah

	catatTransaksi(idPengirim, "Kirim", jumlah, "Transfer ke "+daftarAkun[indexPenerima].Nama)
	catatTransaksi(idPenerima, "Terima", jumlah, "Transfer dari "+daftarAkun[indexPengirim].Nama)

	fmt.Println("Transfer berhasil.")
}

func terimaUang(idAkun string, jumlah float64, keterangan string) {
	index := cariAkun(idAkun)
	if index == -1 {
		fmt.Println("Akun tidak ditemukan.")
		return
	}

	daftarAkun[index].Saldo += jumlah
	catatTransaksi(idAkun, "Terima", jumlah, keterangan)
}

func bayar() {
	var idAkun string
	var jumlah float64
	var keterangan string

	fmt.Print("Masukkan ID Akun Anda: ")
	fmt.Scanln(&idAkun)
	fmt.Print("Masukkan Jumlah Pembayaran: ")
	fmt.Scanln(&jumlah)
	fmt.Print("Masukkan Keterangan Pembayaran: ")
	fmt.Scanln(&keterangan)

	index := cariAkun(idAkun)
	if index == -1 {
		fmt.Println("Akun tidak ditemukan.")
		return
	}

	if daftarAkun[index].Saldo < jumlah {
		fmt.Println("Saldo tidak mencukupi.")
		return
	}

	daftarAkun[index].Saldo -= jumlah
	catatTransaksi(idAkun, "Bayar", jumlah, keterangan)

	fmt.Println("Pembayaran berhasil.")
}

func isiSaldo() {
	var idAkun string
	var jumlah float64

	fmt.Print("Masukkan ID Akun Anda: ")
	fmt.Scanln(&idAkun)
	fmt.Print("Masukkan Jumlah Saldo yang Akan Ditambahkan: ")
	fmt.Scanln(&jumlah)

	index := cariAkun(idAkun)
	if index == -1 {
		fmt.Println("Akun tidak ditemukan.")
		return
	}

	daftarAkun[index].Saldo += jumlah
	catatTransaksi(idAkun, "Isi Saldo", jumlah, "Isi Saldo")

	fmt.Println("Isi saldo berhasil.")
}

func catatTransaksi(akunID string, jenis string, jumlah float64, keterangan string) {
	if jumlahTransaksi >= maxTransaksi {
		fmt.Println("Maaf, kapasitas transaksi penuh.")
		return
	}

	var transaksi Transaksi
	transaksi.ID = buatID()
	transaksi.Jenis = jenis
	transaksi.Jumlah = jumlah
	transaksi.Keterangan = keterangan
	transaksi.AkunID = akunID

	daftarTransaksi[jumlahTransaksi] = transaksi
	jumlahTransaksi++
}

func ubahNamaAkun() {
	var id string
	fmt.Print("Masukkan ID Akun yang namanya akan diubah: ")
	fmt.Scanln(&id)

	index := cariAkun(id)
	if index == -1 {
		fmt.Println("Akun tidak ditemukan.")
		return
	}

	fmt.Printf("Nama akun saat ini: %s\n", daftarAkun[index].Nama)
	fmt.Print("Masukkan Nama baru: ")
	var namaBaru string
	fmt.Scanln(&namaBaru)

	daftarAkun[index].Nama = namaBaru
	fmt.Println("Nama akun berhasil diubah.")

}

func tampilkanRiwayatTransaksi(idAkun string) {
	var riwayat [maxTransaksi]Transaksi
	var jumlahRiwayat int = 0

	for i := 0; i < jumlahTransaksi; i++ {
		if daftarTransaksi[i].AkunID == idAkun {
			if jumlahRiwayat < maxTransaksi {
				riwayat[jumlahRiwayat] = daftarTransaksi[i]
				jumlahRiwayat++
			} else {
				return
			}
		}
	}

	if jumlahRiwayat == 0 {
		fmt.Println("Tidak ada riwayat transaksi untuk akun ini.")
		return
	}

	fmt.Println("Riwayat Transaksi Akun " + idAkun + ":")
	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Printf("%-25s%-10s%-15s%-20s\n", "ID Transaksi", "Jenis", "Jumlah", "Keterangan")
	fmt.Println("--------------------------------------------------------------------------------------------------")
	for i := 0; i < jumlahRiwayat; i++ {
		transaksi := riwayat[i]
		fmt.Printf("%-25s%-10s%-15.2f%-20s\n", transaksi.ID, transaksi.Jenis, transaksi.Jumlah, transaksi.Keterangan)
	}
	fmt.Println("--------------------------------------------------------------------------------------------------")
}

func main() {
	for {
		fmt.Println("\n=== Aplikasi E-Money ===")
		fmt.Println("1. Registrasi Akun")
		fmt.Println("2. Persetujuan/Penolakan Akun (Admin)")
		fmt.Println("3. Kirim Uang")
		fmt.Println("4. Bayar")
		fmt.Println("5. Tampilkan Riwayat Transaksi")
		fmt.Println("6. Tampilkan Daftar Akun ")
		fmt.Println("7. Isi Saldo")
		fmt.Println("8. Urutkan Akun")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			registrasiAkun()
		case 2:
			persetujuanAkun()
		case 3:
			kirimUang()
		case 4:
			bayar()
		case 5:
			var id string
			fmt.Print("Masukkan ID Akun untuk melihat riwayat transaksi: ")
			fmt.Scanln(&id)
			tampilkanRiwayatTransaksi(id)
		case 6:
			tampilkanDaftarAkun()
		case 7:
			isiSaldo()
		case 8:
			fmt.Println("\nPilih arah pengurutan:")
			fmt.Println("1. Ascending (A-Z)")
			fmt.Println("2. Descending (Z-A)")
			fmt.Print("Pilih: ")
			var arahUrut int
			fmt.Scanln(&arahUrut)

			if arahUrut == 1 {
				urutkanAkun(true)
				fmt.Println("\n--- Daftar Akun setelah diurutkan ---")
				fmt.Println("--------------------------------------------------")
				fmt.Printf("%-15s%-20s%-10s%-20s\n", "ID", "Nama", "Saldo", "Status")
				fmt.Println("--------------------------------------------------")
				for i := 0; i < jumlahAkun; i++ {
					fmt.Printf("%-15s%-20s%-10.2f%-20s\n", daftarAkun[i].ID, daftarAkun[i].Nama, daftarAkun[i].Saldo, daftarAkun[i].Status)
				}
				fmt.Println("--------------------------------------------------")
			} else if arahUrut == 2 {
				urutkanAkun(false)
				fmt.Println("\n--- Daftar Akun setelah diurutkan ---")
				fmt.Println("--------------------------------------------------")
				fmt.Printf("%-15s%-20s%-10s%-20s\n", "ID", "Nama", "Saldo", "Status")
				fmt.Println("--------------------------------------------------")
				for i := 0; i < jumlahAkun; i++ {
					fmt.Printf("%-15s%-20s%-10.2f%-20s\n", daftarAkun[i].ID, daftarAkun[i].Nama, daftarAkun[i].Saldo, daftarAkun[i].Status)
				}
				fmt.Println("--------------------------------------------------")
			} else {
				fmt.Println("Pilihan arah pengurutan tidak valid.")
			}
		case 9:
			ubahNamaAkun()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
