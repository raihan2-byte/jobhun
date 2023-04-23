package storage

import "database/sql"

type MahasiswaDao struct {
	Connection *sql.DB
}

func (m *MahasiswaDao) Save(ma *Mahasiswa) (*Mahasiswa, error) {

	stmt, err := m.Connection.Prepare("insert into mahasiswa(nama,usia,gender,tanggal_registrasi) values(?,?,?,?)")

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(ma.Nama, ma.Usia, ma.Gender, ma.TanggalRegistrasi)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	ma.ID = id

	return ma, nil
}

func (m *MahasiswaDao) Find(ID int) (*Mahasiswa, error) {
	stmt, err := m.Connection.Prepare("select id, nama, usia, gender from mahasiswa where id = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(ID)

	if err != nil {
		return nil, err
	}

	j := &Mahasiswa{}

	if rows.Next() {
		rows.Scan(&j.ID, &j.Nama, &j.Usia, &j.Usia, &j.Gender)
		return j, nil
	}

	return nil, nil
}
