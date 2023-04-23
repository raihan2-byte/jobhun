package storage

import "time"

type Mahasiswa struct {
	ID                int64
	Nama              string
	Usia              int
	Gender            int
	TanggalRegistrasi time.Time
}
