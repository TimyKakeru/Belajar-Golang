package Controller

import (
	"Proyek/Model"
	"Proyek/Node"
)

func InputanSiswa(firstName string, lastName string, umur int, kelamin string, dosenWaliId int) bool {
	return Model.InputanSiswa(firstName, lastName, umur, kelamin, dosenWaliId)
}

func MembersView() []Node.Data_Siswa {
	siswa := Model.MemberReadAll()
	if siswa == nil {
		return nil
	}
	return siswa
}

func MemberDelete(id int) (bool, Node.Data_Siswa) {
	member := Node.Data_Siswa{}
	if id < 1 {
		return false, member
	}
	return true, Model.MemberDelete(id)
}

func MemberUpdate(id int, namaDepan string, namaBelakang string, umur int, kelamin string, dosenWaliId int) Node.Data_Siswa {
	return Model.MemberUpdate(id, namaDepan, namaBelakang, umur, kelamin, dosenWaliId)
}

func MemberSearch(id int) Node.Data_Siswa {
	address := Model.MemberSearch(id)
	if address != nil {
		return address.Next.Data
	}
	return Node.Data_Siswa{}
}
func InputanDosen(firstName string, lastName string, umur int, kelamin string, jabatan string) bool {
	if firstName != "" && kelamin != "" && jabatan != "" {
		Model.InputanDosen(firstName, lastName, umur, kelamin, jabatan)
		return true
	}
	return false
}

func DosensView() []Node.Data_Dosen {
	dosen := Model.DosenReadAll()
	if dosen == nil {
		return nil
	}
	return dosen
}

func DosenDelete(id int) (bool, Node.Data_Dosen) {
	dosen := Node.Data_Dosen{}
	if id < 1 {
		return false, dosen
	}
	return true, Model.DosenDelete(id)
}

func DosenUpdate(id int, namaDepan string, namaBelakang string, umur int, kelamin string, jabatan string) Node.Data_Dosen {
	return Model.DosenUpdate(id, namaDepan, namaBelakang, umur, kelamin, jabatan)
}

func DosenSearch(id int) Node.Data_Dosen {
	address := Model.DosenSearch(id)
	if address != nil {
		return address.Next.Data
	}
	return Node.Data_Dosen{}
}
