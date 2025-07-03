package View

import (
	"Proyek/Controller"
	"fmt"
)

func Menu() {
	fmt.Println("================================")
	fmt.Println(" Data Siswa & Dosen ")
	fmt.Println("================================")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Lihat Data")
	fmt.Println("3. Cari Data")
	fmt.Println("4. Update Data")
	fmt.Println("5. Hapus Data")
	fmt.Println("6. Keluar")
}

func Insert() {
	var firstName, lastName, kelamin string
	var umur, dosenWaliId int

	fmt.Println("================================")
	fmt.Println(" Data Siswa ")
	fmt.Println("================================")
	fmt.Print("Masukkan Nama Depan :")
	fmt.Scan(&firstName)
	fmt.Print("Masukkan Nama Belakang :")
	fmt.Scan(&lastName)
	fmt.Print("Masukkan Umur :")
	fmt.Scan(&umur)
	fmt.Print("Jenis Kelamin :")
	fmt.Scan(&kelamin)
	fmt.Print("ID Dosen Wali :")
	fmt.Scan(&dosenWaliId)

	berhasil := Controller.InputanSiswa(firstName, lastName, umur, kelamin, dosenWaliId)
	if berhasil {
		fmt.Println("Data Berhasil Ditambahkan")
	} else {
		fmt.Println("Data Gagal Ditambahkan")
	}
}

func ReadAll() {
	siswa := Controller.MembersView()
	if siswa != nil {
		fmt.Println("================================")
		fmt.Println("SEMUA DATA SISWA ")
		fmt.Println("================================")
		for _, Siswa := range siswa {
			fmt.Println("------------------------------")
			fmt.Println(" Id : ", Siswa.Id)
			fmt.Println(" Nama Depan : ", Siswa.NamaDepan)
			fmt.Println(" Nama Belakang : ", Siswa.NamaBelakang)
			fmt.Println(" Umur : ", Siswa.Umur)
			fmt.Println(" Jenis Kelamin :", Siswa.Kelamin)
			fmt.Println(" Dosen Wali ID: ", Siswa.DosenWali.Id)
			fmt.Println(" Dosen Wali Nama: ", Siswa.DosenWali.NamaDepan, Siswa.DosenWali.NamaBelakang)
		}
	} else {
		fmt.Println("Data Siswa Tidak Ada")
	}
}

func Search() {
	var Id int
	fmt.Print("Inputkan ID Siswa: ")
	fmt.Scan(&Id)
	siswa := Controller.MemberSearch(Id)
	if siswa.Id != 0 {
		fmt.Println("------------------------------")
		fmt.Println(" Id : ", siswa.Id)
		fmt.Println(" Nama Depan : ", siswa.NamaDepan)
		fmt.Println(" Nama Belakang : ", siswa.NamaBelakang)
		fmt.Println(" Umur : ", siswa.Umur)
		fmt.Println(" Jenis Kelamin :", siswa.Kelamin)
		fmt.Println(" Dosen Wali ID: ", siswa.DosenWali.Id)
		fmt.Println(" Dosen Wali Nama: ", siswa.DosenWali.NamaDepan, siswa.DosenWali.NamaBelakang)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func Update() {
	var Id int
	fmt.Print("Masukkan ID Siswa yang mau diperbaharui: ")
	fmt.Scan(&Id)
	siswa := Controller.MemberSearch(Id)
	if siswa.Id == 0 {
		fmt.Println("Data Siswa Tidak Ditemukan")
		return
	}
	fmt.Println("================================")
	fmt.Println("--Detail Data Siswa--")
	fmt.Println("================================")
	fmt.Println(" Id : ", siswa.Id)
	fmt.Println(" Nama Depan : ", siswa.NamaDepan)
	fmt.Println(" Nama Belakang : ", siswa.NamaBelakang)
	fmt.Println(" Umur : ", siswa.Umur)
	fmt.Println(" Jenis Kelamin :", siswa.Kelamin)
	var firstName, lastName, kelamin string
	var umur, dosenWaliId int

	fmt.Println("================================")
	fmt.Println("UPDATE DATA")
	fmt.Println("================================")
	fmt.Print("Masukkan Nama Depan :")
	fmt.Scan(&firstName)
	fmt.Print("Masukkan Nama Belakang :")
	fmt.Scan(&lastName)
	fmt.Print("Masukkan Umur :")
	fmt.Scan(&umur)
	fmt.Println("Jenis Kelamin :")
	fmt.Scan(&kelamin)
	fmt.Print("ID Dosen Wali :")
	fmt.Scan(&dosenWaliId)
	siswaUpdate := Controller.MemberUpdate(Id, firstName, lastName, umur, kelamin, dosenWaliId)
	if siswaUpdate.Id != 0 {
		fmt.Println("Data Berhasil DiUpdate")
		fmt.Println("================================")
		fmt.Println(" Id  :", siswaUpdate.Id)
		fmt.Println(" Nama Depan : ", siswaUpdate.NamaDepan)
		fmt.Println(" Nama Belakang : ", siswaUpdate.NamaBelakang)
		fmt.Println(" Umur : ", siswaUpdate.Umur)
		fmt.Println(" Jenis Kelamin :", siswaUpdate.Kelamin)
		fmt.Println(" Dosen Wali ID: ", siswaUpdate.DosenWali.Id)
		fmt.Println(" Dosen Wali Nama: ", siswaUpdate.DosenWali.NamaDepan, siswaUpdate.DosenWali.NamaBelakang)
	} else {
		fmt.Println("Gagal Mengupdate Data")
	}
}

func Delete() {
	var Id int
	fmt.Print("Masukkan ID Siswa :")
	fmt.Scan(&Id)
	berhasil, _ := Controller.MemberDelete(Id)
	if berhasil {
		fmt.Println("Data berhasil dihapus.")
		fmt.Println("---------------------------")
	} else {
		fmt.Println("Data tidak ditemukan atau terjadi kesalahan saat menghapus data.")
		fmt.Println("---------------------------")
	}
}

func InsertDosen() {
	var firstName, lastName, kelamin, jabatan string
	var umur int

	fmt.Println("================================")
	fmt.Println(" Data Dosen ")
	fmt.Println("================================")
	fmt.Print("Masukkan Nama Depan :")
	fmt.Scan(&firstName)
	fmt.Print("Masukkan Nama Belakang :")
	fmt.Scan(&lastName)
	fmt.Print("Masukkan Umur :")
	fmt.Scan(&umur)
	fmt.Print("Jenis Kelamin :")
	fmt.Scan(&kelamin)
	fmt.Print("Jabatan :")
	fmt.Scan(&jabatan)

	berhasil := Controller.InputanDosen(firstName, lastName, umur, kelamin, jabatan)
	if berhasil {
		fmt.Println("Data Berhasil Ditambahkan")
	} else {
		fmt.Println("Data Gagal Ditambahkan")
	}
}

func ReadAllDosen() {
	dosen := Controller.DosensView()
	if dosen != nil {
		fmt.Println("================================")
		fmt.Println("SEMUA DATA DOSEN ")
		fmt.Println("================================")
		for _, Dosen := range dosen {
			fmt.Println("------------------------------")
			fmt.Println(" Id : ", Dosen.Id)
			fmt.Println(" Nama Depan : ", Dosen.NamaDepan)
			fmt.Println(" Nama Belakang : ", Dosen.NamaBelakang)
			fmt.Println(" Umur : ", Dosen.Umur)
			fmt.Println(" Jenis Kelamin :", Dosen.Kelamin)
			fmt.Println(" Jabatan :", Dosen.Jabatan)
		}
	} else {
		fmt.Println("Data Dosen Tidak Ada")
	}
}

func SearchDosen() {
	var Id int
	fmt.Print("Inputkan ID Dosen: ")
	fmt.Scan(&Id)
	dosen := Controller.DosenSearch(Id)
	if dosen.Id != 0 {
		fmt.Println("------------------------------")
		fmt.Println(" Id : ", dosen.Id)
		fmt.Println(" Nama Depan : ", dosen.NamaDepan)
		fmt.Println(" Nama Belakang : ", dosen.NamaBelakang)
		fmt.Println(" Umur : ", dosen.Umur)
		fmt.Println(" Jenis Kelamin :", dosen.Kelamin)
		fmt.Println(" Jabatan :", dosen.Jabatan)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func UpdateDosen() {
	var Id int
	fmt.Print("Masukkan ID Dosen yang mau diperbaharui: ")
	fmt.Scan(&Id)
	dosen := Controller.DosenSearch(Id)
	if dosen.Id == 0 {
		fmt.Println("Data Dosen Tidak Ditemukan")
		return
	}
	fmt.Println("================================")
	fmt.Println("--Detail Data Dosen--")
	fmt.Println("================================")
	fmt.Println(" Id : ", dosen.Id)
	fmt.Println(" Nama Depan : ", dosen.NamaDepan)
	fmt.Println(" Nama Belakang : ", dosen.NamaBelakang)
	fmt.Println(" Umur : ", dosen.Umur)
	fmt.Println(" Jenis Kelamin :", dosen.Kelamin)
	fmt.Println(" Jabatan :", dosen.Jabatan)
	var firstName, lastName, kelamin, jabatan string
	var umur int

	fmt.Println("================================")
	fmt.Println("UPDATE DATA")
	fmt.Println("================================")
	fmt.Print("Masukkan Nama Depan :")
	fmt.Scan(&firstName)
	fmt.Print("Masukkan Nama Belakang :")
	fmt.Scan(&lastName)
	fmt.Print("Masukkan Umur :")
	fmt.Scan(&umur)
	fmt.Println("Jenis Kelamin :")
	fmt.Scan(&kelamin)
	fmt.Println("Jabatan :")
	fmt.Scan(&jabatan)
	dosenUpdate := Controller.DosenUpdate(Id, firstName, lastName, umur, kelamin, jabatan)
	if dosenUpdate.Id != 0 {
		fmt.Println("Data Berhasil DiUpdate")
		fmt.Println("================================")
		fmt.Println(" Id  :", dosenUpdate.Id)
		fmt.Println(" Nama Depan : ", dosenUpdate.NamaDepan)
		fmt.Println(" Nama Belakang : ", dosenUpdate.NamaBelakang)
		fmt.Println(" Umur : ", dosenUpdate.Umur)
		fmt.Println(" Jenis Kelamin :", dosenUpdate.Kelamin)
		fmt.Println(" Jabatan :", dosenUpdate.Jabatan)
	} else {
		fmt.Println("Gagal Mengupdate Data")
	}
}

func DeleteDosen() {
	var Id int
	fmt.Print("Masukkan ID Dosen :")
	fmt.Scan(&Id)
	berhasil, _ := Controller.DosenDelete(Id)
	if berhasil {
		fmt.Println("Data berhasil dihapus.")
		fmt.Println("---------------------------")
	} else {
		fmt.Println("Data tidak ditemukan atau terjadi kesalahan saat menghapus data.")
		fmt.Println("---------------------------")
	}
}
