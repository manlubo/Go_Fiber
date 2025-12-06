package service

import (
	"fiber/internal/entity"
	"fiber/internal/module/board/repository"
)

type BoardService struct {
	repo *repository.BoardRepository
}

func NewBoardService(repo *repository.BoardRepository) *BoardService {
	return &BoardService{repo: repo}
}

func (s *BoardService) CreateBoard(board *entity.Board) error {
	return s.repo.Create(board)
}

func (s *BoardService) GetBoard(id string) (*entity.Board, error) {
	return s.repo.FindByID(id)
}

func (s *BoardService) GetBoards() ([]*entity.Board, error) {
	return s.repo.List()
}

func (s *BoardService) DeleteBoard(id string) error {
	return s.repo.Delete(id)
}

func (s *BoardService) UpdateBoard(id string, board *entity.Board) error {
	return s.repo.Update(id, board)
}
