package storage

import "database/sql"

type HobiDao struct {
	Connection *sql.DB
}

func (h *HobiDao) Save(hb *Hobi) (*Hobi, error) {
	stmt, err := h.Connection.Prepare("insert into hobi(nama) values(?)")

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

func (h *HobiDao) Find(ID int64) (*Hobi, error) {
	stmt, err := h.Connection.Prepare("select * from hobi where id=?")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}

	j := &Hobi{}

	if rows.Next() {
		err := rows.Scan(&j.ID, &j.Nama)
		if err != nil {
			return nil, err
		}
		return j, nil
	}

	return nil, nil

}
