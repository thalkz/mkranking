package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/thalkz/kart/config"
)

var (
	host            = os.Getenv("POSTGRES_HOST")
	port            = os.Getenv("POSTGRES_PORT")
	user            = os.Getenv("POSTGRES_USER")
	password        = os.Getenv("POSTGRES_PASSWORD")
	dbname          = os.Getenv("POSTGRES_DB")
	migrationFolder = os.Getenv("MIGRATIONS_FOLDER")
)

var db *sql.DB

func Open(cfg *config.Config) (func() error, error) {
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
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	// Create tables if missing
	if err = createPlayersTable(); err != nil {
		return nil, fmt.Errorf("failed to create players table: %w", err)
	}
	if err = createRacesTable(); err != nil {
		return nil, fmt.Errorf("failed to create races table: %w", err)
	}
	if err = createPlayersRacesTable(cfg); err != nil {
		return nil, fmt.Errorf("failed to create players_races table: %w", err)
	}
	if err = createMetadataTable(cfg); err != nil {
		return nil, fmt.Errorf("failed to create metadata table: %w", err)
	}

	if err = applyMigrations(); err != nil {
		return nil, fmt.Errorf("failed to apply migration: %w", err)
	}

	return db.Close, nil
}

// If a migration file with the name of the current version is found, then the migration is applied
// Migrations should update the version in the metadata table to a unique unused name, to prevent cyclic migrations
func applyMigrations() error {
	usedVersions := make(map[string]bool)
	for {
		version, err := getCurrentVersion()
		if err != nil {
			return fmt.Errorf("failed to get current version: %w", err)
		}

		// check cyclic migration
		if usedVersions[version] {
			return fmt.Errorf("cyclic migration detected: database version %v already used", version)
		}
		usedVersions[version] = true

		// get .sql migration
		path := filepath.Join(migrationFolder, version+".sql")
		migrationStr, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			log.Printf("no migration file at: %v\n", path)
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to read migration file %v: %w", path, err)
		}

		// apply migration
		_, err = db.Exec(string(migrationStr))
		if err != nil {
			return fmt.Errorf("failed to apply migration: %w", err)
		}

		log.Printf("Applied migration from file %v\n", path)
	}
}

func getCurrentVersion() (string, error) {
	row := db.QueryRow(`SELECT value FROM metadata WHERE key = 'version'`)
	var version string
	err := row.Scan(&version)
	if err != nil {
		return "", fmt.Errorf("failed to select version metadata: %w", err)
	}
	err = row.Err()
	if err != nil {
		return "", fmt.Errorf("failed to select version metadata: %w", err)
	}
	return version, nil
}

func createPlayersTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS players (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name text NOT NULL,
		rating real NOT NULL,
		icon integer NOT NULL,
		races_count integer NOT NULL DEFAULT 0,
		season integer NOT NULL DEFAULT 1
	);`
	_, err := db.Exec(statement)
	return err
}

func createRacesTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS races (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		ranking integer[] NOT NULL,
		date timestamp without time zone NOT NULL,
		season integer NOT NULL DEFAULT 1
	);`
	_, err := db.Exec(statement)
	return err
}

func createPlayersRacesTable(cfg *config.Config) error {
	statement := `
	CREATE TABLE IF NOT EXISTS players_races (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id integer NOT NULL REFERENCES players(id) ON DELETE CASCADE ON UPDATE CASCADE,
		race_id integer NOT NULL REFERENCES races(id) ON DELETE CASCADE ON UPDATE CASCADE,
		old_rating real NOT NULL DEFAULT %v,
		new_rating real NOT NULL DEFAULT %v,
		rating_diff real NOT NULL DEFAULT 0
	);`
	_, err := db.Exec(fmt.Sprintf(statement, cfg.InitialRating, cfg.InitialRating))
	return err
}

func createMetadataTable(cfg *config.Config) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS metadata (
		key text PRIMARY KEY,
		value text NOT NULL,
		updated_at timestamp without time zone NOT NULL DEFAULT now()
	);`)
	if err != nil {
		return fmt.Errorf("failed to create metadata table: %w", err)
	}

	version, _ := getCurrentVersion()
	if version == "" {
		_, err := db.Exec(`INSERT INTO metadata (key, value) values('version', 'v1')`)
		if err != nil {
			return fmt.Errorf("failed to initialise version to v1: %w", err)
		}
	}

	return nil
}
