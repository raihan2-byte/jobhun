package storage

func (m *MahasiswaDao) Update(ma *Mahasiswa) (*Mahasiswa, error) {

	stmt, err := m.Connection.Prepare("insert into mahasiswa(nama,usia,gender) values(?,?,?)")

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
