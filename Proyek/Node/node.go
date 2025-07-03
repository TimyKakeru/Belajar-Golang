package Node

type Data_Dosen struct {
	Id           int
	NamaDepan    string
	NamaBelakang string
	Umur         int
	Kelamin      string
	Jabatan      string
}

type Data_Siswa struct {
	Id           int
	NamaDepan    string
	NamaBelakang string
	Umur         int
	Kelamin      string
	DosenWali    *Data_Dosen
}

type DataLL struct {
	Data Data_Siswa
	Next *DataLL
}

type DataLLDosen struct {
	Data Data_Dosen
	Next *DataLLDosen
}
