package storage

import "errors"

func (m *MahasiswaDao) Delete(ma *Mahasiswa) error {
	stmt, err := m.Connection.Prepare("delete from mahasiswa where id=?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(ma.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
