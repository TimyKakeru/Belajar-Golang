package Model

import (
	"Proyek/Database"
	"Proyek/Node"
)

func InputanSiswa(firstName string, lastName string, umur int, kelamin string, dosenWaliId int) bool {
	tempLL := &Database.DBMember

	dosenWali := MemberSearchDosen(dosenWaliId)
	if dosenWali == nil {
		return false
	}

	member := Node.Data_Siswa{
		Id:           memberLastId(),
		NamaDepan:    firstName,
		NamaBelakang: lastName,
		Umur:         umur,
		Kelamin:      kelamin,
		DosenWali:    &dosenWali.Data,
	}

	newLL := Node.DataLL{
		Data: member,
	}
	if tempLL.Next == nil {
		tempLL.Next = &newLL
	} else {
		current := tempLL
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newLL
	}
	return true
}

func memberLastId() int {
	tempLL := &Database.DBMember

	if tempLL.Next == nil {
		return 1
	}

	current := tempLL
	for current.Next != nil {
		current = current.Next
	}
	return current.Data.Id + 1
}

func MemberReadAll() []Node.Data_Siswa {
	tempLL := &Database.DBMember
	var memberTable []Node.Data_Siswa

	current := tempLL
	for current.Next != nil {
		current = current.Next
		memberTable = append(memberTable, current.Data)
	}
	return memberTable
}

func MemberSearch(id int) *Node.DataLL {
	tempLL := &Database.DBMember

	if tempLL.Next != nil {
		current := tempLL
		for current.Next != nil {
			if current.Next.Data.Id == id {
				return current
			}
			current = current.Next
		}
	}
	return nil
}

func MemberDelete(id int) Node.Data_Siswa {
	address := MemberSearch(id)
	if address != nil {
		nodeDelete := address.Next
		address.Next = address.Next.Next
		return nodeDelete.Data
	}
	return Node.Data_Siswa{}
}

func MemberUpdate(id int, namaDepan string, namaBelakang string, umur int, kelamin string, dosenWaliId int) Node.Data_Siswa {
	address := MemberSearch(id)
	if address != nil {
		address.Next.Data.NamaDepan = namaDepan
		address.Next.Data.NamaBelakang = namaBelakang
		address.Next.Data.Umur = umur
		address.Next.Data.Kelamin = kelamin

		dosenWali := MemberSearchDosen(dosenWaliId)
		if dosenWali != nil {
			address.Next.Data.DosenWali = &dosenWali.Data
		}

		return address.Next.Data
	}
	return Node.Data_Siswa{}
}

func MemberSearchDosen(id int) *Node.DataLLDosen {
	tempLL := &Database.DBDosen

	if tempLL.Next != nil {
		current := tempLL
		for current.Next != nil {
			if current.Next.Data.Id == id {
				return current.Next
			}
			current = current.Next
		}
	}
	return nil
}

func InputanDosen(firstName string, lastName string, umur int, kelamin string, jabatan string) {
	tempLL := &Database.DBDosen

	dosen := Node.Data_Dosen{
		Id:           dosenLastId(),
		NamaDepan:    firstName,
		NamaBelakang: lastName,
		Umur:         umur,
		Kelamin:      kelamin,
		Jabatan:      jabatan,
	}

	newLL := Node.DataLLDosen{
		Data: dosen,
	}
	if tempLL.Next == nil {
		tempLL.Next = &newLL
	} else {
		current := tempLL
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newLL
	}
}

func dosenLastId() int {
	tempLL := &Database.DBDosen

	if tempLL.Next == nil {
		return 1
	}

	current := tempLL
	for current.Next != nil {
		current = current.Next
	}
	return current.Data.Id + 1
}

func DosenReadAll() []Node.Data_Dosen {
	tempLL := &Database.DBDosen
	var dosenTable []Node.Data_Dosen

	current := tempLL
	for current.Next != nil {
		current = current.Next
		dosenTable = append(dosenTable, current.Data)
	}
	return dosenTable
}

func DosenSearch(id int) *Node.DataLLDosen {
	tempLL := &Database.DBDosen

	if tempLL.Next != nil {
		current := tempLL
		for current.Next != nil {
			if current.Next.Data.Id == id {
				return current
			}
			current = current.Next
		}
	}
	return nil
}

func DosenDelete(id int) Node.Data_Dosen {
	address := DosenSearch(id)
	if address != nil {
		nodeDelete := address.Next
		address.Next = address.Next.Next
		return nodeDelete.Data
	}
	return Node.Data_Dosen{}
}

func DosenUpdate(id int, namaDepan string, namaBelakang string, umur int, kelamin string, jabatan string) Node.Data_Dosen {
	address := DosenSearch(id)
	if address != nil {
		address.Next.Data.NamaDepan = namaDepan
		address.Next.Data.NamaBelakang = namaBelakang
		address.Next.Data.Umur = umur
		address.Next.Data.Kelamin = kelamin
		address.Next.Data.Jabatan = jabatan
		return address.Next.Data
	}
	return Node.Data_Dosen{}
}
