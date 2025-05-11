package repository

import (
	"database/sql"
	"errors"

	"github.com/buffalo-big-tech-system/bookcenter-grpc-go/internal/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repos *Repository) getPagesCount(countPerPage int) (int, error) {
	stmt, err := repos.db.Prepare("SELECT CEIL(COUNT(*) / ?) AS total_pages FROM Books")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(countPerPage)
	var pagesCount int

	err = row.Scan(&pagesCount)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: add own err
			return 0, err
		}
		return 0, err
	}

	return pagesCount, nil
}

func (repos *Repository) getPage(pageIndex int, pageSize int) ([]model.Book, error) {
	offset := (pageIndex - 1) * pageSize

	stmt, err := repos.db.Prepare("SELECT id, title FROM Books ORDER BY id LIMIT ? OFFSET ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(stmt, pageSize, offset)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: add own err
			return nil, err
		}
		return nil, err
	}

	var books []model.Book
	for rows.Next() {
		var b model.Book

		if err = rows.Scan(&b.Id, &b.Title); err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}
