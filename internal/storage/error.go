package storage

import "errors"

var (
	ErrBooksNotFound = errors.New("Books are not found")
	ErrPagesNotFound = errors.New("Pages are not found")
)
