package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Storage interface {
	GetConnection() *sql.DB
}

type MysqlStorage struct {
}

func (m *MysqlStorage) GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/jobhun")
	if err != nil {
		fmt.Printf("an error occurred %v", err)
		return nil
	}

	return db
}
