package service

import (
	"fmt"
	"log"

	"github.com/buffalo-big-tech-system/bookcenter-grpc-go/internal/model"
	"github.com/buffalo-big-tech-system/bookcenter-grpc-go/internal/storage/repository"
)

type Service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}

func (s *Service) GetPagesCount(pagesSize int) (int, error) {
	const op = "service.GetPagesCount"
	log.Printf("Start %s", op)

	count, err := s.repos.GetPagesCount(pagesSize)

	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("Finish %s", op)
	return count, nil
}

func (s *Service) GetPage(pageIndex int, pageSize int) ([]model.Book, error) {
	const op = "service.GetPagesCount"
	log.Printf("Start %s", op)

	books, err := s.repos.GetPage(pageIndex, pageSize)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("Finish %s", op)
	return books, nil
}
