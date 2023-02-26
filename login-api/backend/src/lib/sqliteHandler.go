package lib

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) AddNewUser(id string, password string) *User {
	sql, err := s.db.Prepare("INSERT INTO users (id, password, createdAt) VALUES (?, ?, datetime('now'))")
	if err != nil {
		panic(err)
	}
	_, err = sql.Exec(id, password)
	if err != nil {
		panic(err)
	}
	var user User
	user.ID = id
	user.Password = password
	user.CreateAt = time.Now()
	return &user

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
