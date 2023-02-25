package lib

import "time"

type User struct {
	Seq      int       `json:"seq"`
	ID       string    `json:"id"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"created_at"`
}

type DBHandler interface {
	Close()
}

func NewDBHandler(filepath string) DBHandler {
	return NewSqliteHandler(filepath)
}
