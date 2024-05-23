package postgresql

import (
	"cityger-be/internal/storage/postgresql/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Close() {
	s.db.Close()
}

func New(dbConnect string) (*Storage, error) {
	const fn = "storage.postgresql.New"

	db, err := sql.Open("postgres", dbConnect)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	// TODO: Check DB is initialized to the desired migration level

	return &Storage{db: db}, nil
}

//! FIXME: CRUDs as separate parts of the package (probably)

// TODO: CRUD: Role
func (s *Storage) GetAllRoles() ([]models.Role, error) {
	const fn = "storage.postgresql.GetAllRoles"

	roles, err := models.GetAllRoles(s.db)
	if err != nil {
		return []models.Role{}, fmt.Errorf("%s: %w", fn, err)
	}
	return roles, nil
}

func (s *Storage) AddNewRole(name, info string) (int, error) {
	const fn = "storage.postgresql.AddNewRole"

	id, err := models.AddRole(s.db, name, info)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", fn, err)
	}

	return id, nil
}

// TODO: CRUD: Company

// TODO: CRUD: User

// TODO: Add user
func (s *Storage) AddUser() {

}

// TODO: Get user

// TODO: Get users for Company

// TODO: Update user

// TODO: Delete user (by id)

// TODO: CRUD: Project

// TODO: CRUD: Scenario

// TODO: CRUD: Zones & Cells
