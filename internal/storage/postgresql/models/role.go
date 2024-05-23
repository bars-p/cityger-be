package models

import (
	"cityger-be/internal/storage"
	"database/sql"

	"github.com/lib/pq"
)

type Role struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Info   string `db:"info"`
	Active bool   `db:"active"`
}

// TODO: Is get single row needed at all?
// func GetRole(db *sql.DB, id int) (Role, error) {
// 	return Role{}, nil
// }

func GetAllRoles(db *sql.DB) ([]Role, error) {
	rows, err := db.Query(`SELECT * FROM "role"`)
	if err != nil {
		return []Role{}, err
	}
	defer rows.Close()

	var roles []Role

	for rows.Next() {
		var r Role
		err := rows.Scan(&r.ID, &r.Name, &r.Info, &r.Active)
		if err != nil {
			return []Role{}, err
		}
		roles = append(roles, r)
	}
	if err = rows.Err(); err != nil {
		return []Role{}, err
	}

	return roles, nil
}

func AddRole(db *sql.DB, name, info string) (int, error) {
	var id int
	err := db.QueryRow(`INSERT INTO "role"(name, info) VALUES($1, $2) RETURNING id`, name, info).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
			return -1, storage.ErrDuplicateRoleName
		}
		return -1, err
	}

	//* Other way to obtain the new row ID value due to lack of LastInsertId support.
	//* On error return 0 meaning that the new value was stored but its ID is unknown.
	// var id int
	// err = db.QueryRow(`SELECT currval(pg_get_serial_sequence('role','id'))`).Scan(&id)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to get ID value: %w", err)
	// }

	return id, nil
}

func UpdateRole() {}

func DeleteRole(id int) {

}
