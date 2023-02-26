package lib

import "time"

type User struct {
	ID       string    `json:"id"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"created_at"`
}

type DBHandler interface {
	AddNewUser(id string, password string) *User
	Close()
}

func NewDBHandler(filepath string) DBHandler {
	return NewSqliteHandler(filepath)
}
