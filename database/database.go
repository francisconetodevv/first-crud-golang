package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Connection drive to the Mysql
)

// Creating the connecting with the Database
func Connection() (*sql.DB, error) {
	stringConnection := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// Connection with the database
	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
