package sqlite

import "database/sql"

type Sqlite struct {
	Database *sql.DB
	Name     string
}

func New(Name string) (*Sqlite, error) {

	s := new(Sqlite)

	database, err := sql.Open("sqlite3", Name)
	if err != nil {
		return nil, err
	}

	s.Database = database
	s.Name = Name

	statement := `
	DROP TABLE IF EXISTS info;
	CREATE TABLE info (id INTEGER PRIMARY KEY, name TEXT, day INTEGER, month INTEGER);
	INSERT INTO info(name, day, month) VALUES('Arthur', 14, 9);
	INSERT INTO info(name, day, month) VALUES('Percival', 14, 9);
	INSERT INTO info(name, day, month) VALUES('Gawain', 14, 9);
	INSERT INTO info(name, day, month) VALUES('Lancelot', 14, 9);
	INSERT INTO info(name, day, month) VALUES('Tristan', 14, 9);
	DROP TABLE IF EXISTS channels;
	CREATE TABLE channels (id INTEGER PRIMARY KEY, telegramId INTEGER);
	INSERT INTO channels(telegramId) VALUES(-1001985611485);
	INSERT INTO channels(telegramId) VALUES(-1001863618407);`

	_, err = s.Database.Exec(statement)
	if err != nil {
		return nil, err
	}

	return s, nil
}
