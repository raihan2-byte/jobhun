package storage

import "database/sql"

type MahasiswaJurusanDao struct {
	Connection *sql.DB
}

func (h *MahasiswaJurusanDao) Save(hb *MahasiwaJurusan) (*MahasiwaJurusan, error) {
	stmt, err := h.Connection.Prepare("insert into mahasiswa_jurusan(id_mahasiswa,id_jurusan) values(?,?)")

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(hb.IDMahasiswa, hb.IDJurusan)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	hb.ID = id

	return hb, nil
}

func (h *MahasiswaJurusanDao) Find(idMahasiswa int) (*MahasiwaJurusan, error) {
	stmt, err := h.Connection.Prepare("select id, id_mahasiswa, id_jurusan from mahasiswa_jurusan where id_mahasiswa = ?")

	if err != nil {
		return nil, err
	}

	//exec manipulasi data & query ga ada data yg diubah

	res, err := stmt.Query(idMahasiswa)

	if err != nil {
		return nil, err
	}

	j := &MahasiwaJurusan{}

	if res.Next() {
		res.Scan(&j.ID, &j.IDMahasiswa, &j.IDJurusan)
		return j, nil
	}

	return nil, nil

}
