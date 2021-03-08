package db

import (
	"ansme/src/config"
	"ansme/src/utils/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	dbConfig := config.DBConfig()

	driver := dbConfig.Driver

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	db, error := sql.Open(driver, DSN)
	if error != nil {
		logger.Error(fmt.Sprintf("connection error: %s", error.Error()))
	}

	return db, error
}

func QueryRows(sql string) (*sql.Rows, error) {
	db, error := connect()
	if error != nil {
		logger.Error(error.Error())
		return nil, error
	}

	defer db.Close()

	rows, error := db.Query(sql)
	if error != nil {
		logger.Error(error.Error())
		return nil, error
	}

	return rows, nil
}

func QueryOne(sql string) (*sql.Row, error) {
	db, error := connect()
	if error != nil {
		logger.Error(error.Error())
		return nil, error
	}

	defer db.Close()

	row := db.QueryRow(sql)

	return row, error
}

func Exectue(sql string) bool {

	db, error := connect()
	if error != nil {
		logger.Error(error.Error())
		return false
	}
	defer db.Close()

	_, error = db.Exec(sql)
	if error != nil{
		logger.Error(error.Error())
		return false
	}

	return true
}
