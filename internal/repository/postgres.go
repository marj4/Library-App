package repository

import (
	"database/sql"
	"library/error2"
)

const (
	employee = "employee"
)

func Connect(ConnectString string) (*sql.DB, error) {
	//Create connect
	db, err := sql.Open("postgres", ConnectString)
	if err != nil {
		return nil, error2.ErrorWrap("Cant create connect to db:", err)
	}

	//Check connect
	if err = db.Ping(); err != nil {
		return nil, error2.ErrorWrap("Cant ping connect to db:", err)
	}

	//Возвращаем БД
	return db, nil
}
