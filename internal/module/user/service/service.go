package service

import (
	"fiber/internal/entity"
	"fiber/internal/module/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *entity.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUsers() ([]*entity.User, error) {
	return s.repo.List()
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) UpdateUser(id string, user *entity.User) error {
	return s.repo.Update(id, user)
}
