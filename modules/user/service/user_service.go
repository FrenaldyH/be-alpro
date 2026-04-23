package service

import (
	"errors"

	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/Mobilizes/materi-be-alpro/modules/user/dto"
	"github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	"github.com/Mobilizes/materi-be-alpro/pkg/helpers"
	"gorm.io/gorm"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*entities.User, error) {
    hashedPassword, err := helpers.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := &entities.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
    }

    err = s.repo.Create(user)
    return user, err
}

func (s *UserService) GetAllUser() ([]entities.User, error) {
    users, err := s.repo.FindAll()

    return users, err
}

func (s *UserService) GetUserByID(ID uint) (*entities.User, error) {
    user, err := s.repo.FindByID(ID)

    // Development: ---
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("User tidak ditemukan | service")
    }
    return user, err
}

