package lib

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetUsers() []*User {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	users := []*User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Password, &user.CreateAt)
		users = append(users, &user)
	}
	return users

}

// TODO: DB에 유저 안들어가는 오류
func (s *sqliteHandler) AddNewUser(id string, password string) *User {
	stmt, err := s.db.Prepare("INSERT INTO users (id, password, createdAt) VALUES (?, ?, datetime('now'))")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id, password)
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

	stmt, _ := database.Prepare(`
	CREATE TABLE IF NOT EXISTS users (
		id      TEXT,
		password TEXT,
		createdAt DATETIME
	)`)

	stmt.Exec()

	return &sqliteHandler{db: database}
}
