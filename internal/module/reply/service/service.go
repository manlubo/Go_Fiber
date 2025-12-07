package service

import (
	"fiber/internal/entity"
	"fiber/internal/module/reply/repository"
)

type ReplyService struct {
	repo *repository.ReplyRepository
}

func NewReplyService(repo *repository.ReplyRepository) *ReplyService {
	return &ReplyService{repo: repo}
}

func (s *ReplyService) CreateReply(reply *entity.Reply) error {
	return s.repo.Create(reply)
}

func (s *ReplyService) GetReply(id string) (*entity.Reply, error) {
	return s.repo.FindByID(id)
}

func (s *ReplyService) GetReplies() ([]*entity.Reply, error) {
	return s.repo.List()
}

func (s *ReplyService) DeleteReply(id string) error {
	return s.repo.Delete(id)
}

func (s *ReplyService) UpdateReply(id string, reply *entity.Reply) error {
	return s.repo.Update(id, reply)
}
