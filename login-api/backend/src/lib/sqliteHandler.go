package lib

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetUsers() []*User {
	rows, err := s.db.Query("SELECT * FROM users")
	errorHandler(err)
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
	errorHandler(err)
	_, err = stmt.Exec(id, password)
	errorHandler(err)
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
	errorHandler(err)

	stmt, _ := database.Prepare(`
	CREATE TABLE IF NOT EXISTS users (
		id      TEXT PRIMARY KEY,
		password TEXT,
		createdAt DATETIME
	)`)

	_, err = stmt.Exec()

	errorHandler(err)

	return &sqliteHandler{db: database}
}

func errorHandler(err error) {
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
}
