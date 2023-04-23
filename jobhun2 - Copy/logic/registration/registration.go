package registration

import (
	"database/sql"
	"fmt"
	"jobhun2/dto"
	"jobhun2/storage"
	"time"
)

type RegistrationManager struct {
	Connection *sql.DB
}

func (r *RegistrationManager) Init(stg storage.Storage) {
	r.Connection = stg.GetConnection()
}

func (r *RegistrationManager) Terminate(stg storage.Storage) {
	stg.GetConnection().Close()
}

//func (rm *RegistrationManager) Delete(r *dto.CreateMahasiswaRequest) (*dto.CreateMahasiswaResponse, error) {
//
//}

func (rm *RegistrationManager) Update(r *dto.CreateMahasiswaRequest) (*dto.CreateMahasiswaResponse, error) {
	tx, err := rm.Connection.Begin()
	//get data mahasiswa berdasarkan id
	mDao := storage.MahasiswaDao{Connection: rm.Connection}
	mhasiswa, err := mDao.Find(r.ID)
	if err != nil {
		return nil, err
	}

	if mhasiswa == nil {
		return nil, fmt.Errorf("could not find data")
	}

	//get data jurusan mahasiswa berdasarkan id
	mJDao := storage.MahasiswaJurusanDao{Connection: rm.Connection}

	jurusan, err := mJDao.Find(r.ID)
	if err != nil {
		return nil, err
	}

	if jurusan == nil {
		return nil, fmt.Errorf("coul not find data")
	}

	jur := storage.JurusanDao{Connection: rm.Connection}

	mJur, err := jur.Find(jurusan.IDJurusan)
	if err != nil {
		return nil, err
	}

	//get data hobi mahasiswa berdasarkan id
	mHDao := storage.MahasiswaHobiDao{Connection: rm.Connection}

	hobi, err := mHDao.Find(r.ID)
	if err != nil {
		return nil, err
	}

	if hobi == nil {
		return nil, fmt.Errorf("couldd not find data")
	}

	hob := storage.HobiDao{Connection: rm.Connection}

	mHob, err := hob.Find(hobi.IDHobi)
	if err != nil {
		return nil, err
	}

	mhasiswa, err = mDao.Save(&storage.Mahasiswa{
		Nama:   r.Nama,
		Usia:   r.Usia,
		Gender: r.Gender,
	})

	_, err = mHDao.Save(&storage.MahasiwaHobi{
		IDMahasiswa: mhasiswa.ID,
		IDHobi:      r.Hobi,
	})

	_, err = mJDao.Save(&storage.MahasiwaJurusan{
		IDMahasiswa: mhasiswa.ID,
		IDJurusan:   r.Jurusan,
	})

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &dto.CreateMahasiswaResponse{
		Nama:    r.Nama,
		Usia:    r.Usia,
		Gender:  r.Gender,
		Jurusan: mJur,
		Hobi:    mHob,
	}, nil
}

func (rm *RegistrationManager) Register(r *dto.CreateMahasiswaRequest) (*dto.CreateMahasiswaResponse, error) {
	tx, err := rm.Connection.Begin()
	if err != nil {
		return nil, err
	}

	hDao := storage.HobiDao{Connection: rm.Connection}
	jDao := storage.JurusanDao{Connection: rm.Connection}

	hobi, err := hDao.Find(r.Hobi)
	if err != nil || hobi == nil {
		return nil, fmt.Errorf("failed to validate hobby %v", err)
	}

	jurusan, err := jDao.Find(r.Hobi)
	if err != nil || jurusan == nil {
		return nil, fmt.Errorf("failed to validate jurusan %v", err)
	}

	mDao := storage.MahasiswaDao{Connection: rm.Connection}
	mHDao := storage.MahasiswaHobiDao{Connection: rm.Connection}
	mJDao := storage.MahasiswaJurusanDao{Connection: rm.Connection}

	mhasiswa, err := mDao.Save(&storage.Mahasiswa{
		Nama:              r.Nama,
		Usia:              r.Usia,
		Gender:            r.Gender,
		TanggalRegistrasi: time.Now(),
	})

	_, err = mHDao.Save(&storage.MahasiwaHobi{
		IDMahasiswa: mhasiswa.ID,
		IDHobi:      r.Hobi,
	})

	_, err = mJDao.Save(&storage.MahasiwaJurusan{
		IDMahasiswa: mhasiswa.ID,
		IDJurusan:   r.Jurusan,
	})

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &dto.CreateMahasiswaResponse{
		Nama:    r.Nama,
		Usia:    r.Usia,
		Gender:  r.Gender,
		Hobi:    hobi,
		Jurusan: jurusan,
	}, nil
}
