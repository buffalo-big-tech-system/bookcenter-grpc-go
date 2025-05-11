package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/buffalo-big-tech-system/bookcenter-grpc-go/internal/model"
	"github.com/buffalo-big-tech-system/bookcenter-grpc-go/internal/storage"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repos *Repository) GetPagesCount(countPerPage int) (int, error) {
	const op = "repository.GetPagesCount"
	log.Printf("Start %s", op)

	stmt, err := repos.db.Prepare("SELECT CEIL(COUNT(*) / ?) AS total_pages FROM Books")

	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	log.Printf("Query to DB %s", op)

	row := stmt.QueryRow(countPerPage)

	var pagesCount int
	err = row.Scan(&pagesCount)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, fmt.Errorf("%s: %w", op, storage.ErrPagesNotFound)
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("Finish %s", op)
	return pagesCount, nil
}

func (repos *Repository) GetPage(pageIndex int, pageSize int) ([]model.Book, error) {
	const op = "repository.GetPage"
	log.Printf("Start %s", op)

	offset := (pageIndex - 1) * pageSize

	stmt, err := repos.db.Prepare("SELECT id, title FROM Books ORDER BY id LIMIT ? OFFSET ?")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	log.Printf("Query to DB %s", op)
	rows, err := stmt.Query(stmt, pageSize, offset)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrBooksNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var books []model.Book
	for rows.Next() {
		var b model.Book

		if err = rows.Scan(&b.Id, &b.Title); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		books = append(books, b)
	}

	log.Printf("Finish %s", op)
	return books, nil
}
