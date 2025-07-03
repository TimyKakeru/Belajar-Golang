package main

import (
	"Proyek/View"
	"fmt"
)

func main() {
	var pilih, subPilih int
	for {
		View.Menu()
		fmt.Print("Masukkan Pilihan Anda : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			fmt.Println("================================")
			fmt.Println(" Tambah Data ")
			fmt.Println("================================")
			fmt.Println("1. Tambah Data Siswa")
			fmt.Println("2. Tambah Data Dosen")
			fmt.Println("================================")
			fmt.Print("Masukkan Pilihan Anda : ")
			fmt.Scan(&subPilih)
			switch subPilih {
			case 1:
				View.Insert()
			case 2:
				View.InsertDosen()
			default:
				fmt.Println("Pilihan Yang Anda Masukkan Salah")
			}
		case 2:
			fmt.Println("================================")
			fmt.Println(" Lihat Data ")
			fmt.Println("================================")
			fmt.Println("1. Lihat Data Siswa")
			fmt.Println("2. Lihat Data Dosen")
			fmt.Println("================================")
			fmt.Print("Masukkan Pilihan Anda : ")
			fmt.Scan(&subPilih)
			switch subPilih {
			case 1:
				View.ReadAll()
			case 2:
				View.ReadAllDosen()
			default:
				fmt.Println("Pilihan Yang Anda Masukkan Salah")
			}
		case 3:
			fmt.Println("================================")
			fmt.Println(" Cari Data ")
			fmt.Println("================================")
			fmt.Println("1. Cari Data Siswa")
			fmt.Println("2. Cari Data Dosen")
			fmt.Println("================================")
			fmt.Print("Masukkan Pilihan Anda : ")
			fmt.Scan(&subPilih)
			switch subPilih {
			case 1:
				View.Search()
			case 2:
				View.SearchDosen()
			default:
				fmt.Println("Pilihan Yang Anda Masukkan Salah")
			}
		case 4:
			fmt.Println("================================")
			fmt.Println(" Update Data ")
			fmt.Println("================================")
			fmt.Println("1. Update Data Siswa")
			fmt.Println("2. Update Data Dosen")
			fmt.Println("================================")
			fmt.Print("Masukkan Pilihan Anda : ")
			fmt.Scan(&subPilih)
			switch subPilih {
			case 1:
				View.Update()
			case 2:
				View.UpdateDosen()
			default:
				fmt.Println("Pilihan Yang Anda Masukkan Salah")
			}
		case 5:
			fmt.Println("================================")
			fmt.Println(" Hapus Data ")
			fmt.Println("================================")
			fmt.Println("1. Hapus Data Siswa")
			fmt.Println("2. Hapus Data Dosen")
			fmt.Println("================================")
			fmt.Print("Masukkan Pilihan Anda : ")
			fmt.Scan(&subPilih)
			switch subPilih {
			case 1:
				View.Delete()
			case 2:
				View.DeleteDosen()
			default:
				fmt.Println("Pilihan Yang Anda Masukkan Salah")
			}
		case 6:
			fmt.Println("Anda Telah Keluar Dari Program")
			return
		default:
			fmt.Println("Pilihan Yang Anda Masukkan Salah")
		}
		fmt.Println("\nTekan 'Enter' untuk kembali ke menu utama...")
		fmt.Scanln()
		fmt.Scanln()
	}
}
