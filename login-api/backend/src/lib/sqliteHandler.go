package lib

import "database/sql"

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func NewSqliteHandler(filePath string) DBHandler {
	database, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		seq        INTEGER  PRIMARY KEY AUTOINCREMENT,
		id      TEXT,
		password TEXT,
		createdAt DATETIME
	)`
	sql, _ := database.Prepare(createTableQuery)

	sql.Exec()

	return &sqliteHandler{db: database}
}
