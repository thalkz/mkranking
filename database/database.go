package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "kart_test"
)

var db *sql.DB

func Open() error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Printf("Opening database host=%s port=%d dbname=%s...\n", host, port, dbname)
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Database is opened")
	return nil
}

func ExecRaw(statement string) (sql.Result, error) {
	fmt.Printf("[Database] %s\n", statement)
	result, err := db.Exec(statement)
	return result, err
}

func Close() {
	db.Close()
	fmt.Println("Database is closed")
}
