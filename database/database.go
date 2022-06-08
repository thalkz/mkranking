package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

var db *sql.DB

func Open() (func() error, error) {
	// Default config variables
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}

	// Open database connection
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Printf("Opening database host=%s port=%s dbname=%s\n", host, port, dbname)
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect %w", err)
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping %w", err)
	}

	// Create tables if missing
	if err = createPlayersTable(); err != nil {
		return nil, fmt.Errorf("failed to create players table %w", err)
	}
	if err = createRacesTable(); err != nil {
		return nil, fmt.Errorf("failed to create races table %w", err)
	}
	if err = createPlayersRacesTable(); err != nil {
		return nil, fmt.Errorf("failed to create players_races table %w", err)
	}

	return db.Close, nil
}

func createPlayersTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS players (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name text NOT NULL,
		rating real NOT NULL,
		icon integer NOT NULL,
		races_count integer NOT NULL DEFAULT 0
	);`
	_, err := db.Exec(statement)
	return err
}

func createRacesTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS races (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		ranking integer[] NOT NULL,
		date timestamp without time zone NOT NULL
	);`
	_, err := db.Exec(statement)
	return err
}

func createPlayersRacesTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS players_races (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id integer NOT NULL REFERENCES players(id) ON DELETE CASCADE ON UPDATE CASCADE,
		race_id integer NOT NULL REFERENCES races(id) ON DELETE CASCADE ON UPDATE CASCADE,
		old_rating real NOT NULL DEFAULT 1000,
		new_rating real NOT NULL DEFAULT 1000,
		rating_diff real NOT NULL DEFAULT 0
	);`
	_, err := db.Exec(statement)
	return err
}
