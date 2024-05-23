package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int       `db:"id"`
	Login    string    `db:"login"`
	Name     string    `db:"name"`
	Active   bool      `db:"active"`
	Updated  time.Time `db:"updated"` //! Check correct type to store timestamp
	ByUserID int       `db:"by_user_id"`
	ClientID int       `db:"client_id"`
	RoleID   int       `db:"role_id"`
}

func GetUser(db *sql.DB, id int) (User, error) {
	return User{}, nil
}

func GetAllUsers(db *sql.DB, clientID int) ([]User, error) {
	return []User{{}}, nil
}
