package storage

import "database/sql"

type JurusanDao struct {
	Connection *sql.DB
}

func (h *JurusanDao) Save(hb *Jurusan) (*Jurusan, error) {
	stmt, err := h.Connection.Prepare("insert into jurusan(nama) values(?)")

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(hb.Nama)

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

func (h *JurusanDao) Find(ID int64) (*Jurusan, error) {
	stmt, err := h.Connection.Prepare("select * from jurusan where id=?")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}

	j := &Jurusan{}

	if rows.Next() {
		rows.Scan(&j.ID, &j.Nama)
		return j, nil
	}

	return nil, nil

}
