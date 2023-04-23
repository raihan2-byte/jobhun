package storage

import "database/sql"

type MahasiswaHobiDao struct {
	Connection *sql.DB
}

func (h *MahasiswaHobiDao) Save(hb *MahasiwaHobi) (*MahasiwaHobi, error) {
	stmt, err := h.Connection.Prepare("insert into mahasiswa_hobi(id_mahasiswa,id_hobi) values(?,?)")

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(hb.IDMahasiswa, hb.IDHobi)

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

func (h *MahasiswaHobiDao) Find(idMahasiswa int) (*MahasiwaHobi, error) {
	stmt, err := h.Connection.Prepare("select id, id_mahasiswa, id_hobi from mahasiswa_hobi where id_mahasiswa = ?")

	if err != nil {
		return nil, err
	}

	//exec manipulasi data & query ga ada data yg diubah

	res, err := stmt.Query(idMahasiswa)

	if err != nil {
		return nil, err
	}

	j := &MahasiwaHobi{}

	if res.Next() {
		res.Scan(&j.ID, &j.IDMahasiswa, &j.IDHobi)
		return j, nil
	}

	return nil, nil

}
