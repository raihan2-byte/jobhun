package dto

import "jobhun2/storage"

type CreateMahasiswaRequest struct {
	ID      int
	Nama    string
	Usia    int
	Gender  int
	Hobi    int64
	Jurusan int64
}

type CreateMahasiswaResponse struct {
	Nama    string
	Usia    int
	Gender  int
	Hobi    *storage.Hobi
	Jurusan *storage.Jurusan
}
