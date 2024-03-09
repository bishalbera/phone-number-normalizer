package db

import (
	"database/sql"
	
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Phone struct {
	ID int
	Number string
}

type DB struct {
	db *sql.DB
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil

}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE" + name)
	if err != nil {
		return err
	}
	return nil
}