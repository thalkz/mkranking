package database

func CreatePlayersTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS players (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name text,
		rating double precision
	);`
	_, err := db.Exec(statement)
	return err
}

func CreateRacesTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS races (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		ranking integer[]
	);`
	_, err := db.Exec(statement)
	return err
}

func CreatePlayersRacesTable() error {
	statement := `
	CREATE TABLE IF NOT EXISTS players_races (
		id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id integer NOT NULL REFERENCES players(id),
		race_id integer NOT NULL REFERENCES races(id)
	);`
	_, err := db.Exec(statement)
	return err
}
