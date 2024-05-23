package storage

import "errors"

var (
	ErrDataNotFound      = errors.New("data not found in the storage")
	ErrDuplicateRoleName = errors.New("specified NAME for the role is NOT UNIQUE")
)
