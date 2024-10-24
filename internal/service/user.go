package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository interfaces.UserRepository
}

type UserService interface {
	Get(ctx context.Context) (*[]model.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, payload *model.User) (data *model.UserResponse, err error)
	Update(ctx context.Context, payload *model.User, id uuid.UUID) (data *model.UserResponse, err error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewUserService(r *repository.Contract) UserService {
	return &userService{
		UserRepository: r.User,
	}
}

func (s *userService) Get(ctx context.Context) (*[]model.User, error) {
	return s.UserRepository.Get(ctx)
}

func (s *userService) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.UserRepository.GetById(ctx, id)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.UserRepository.GetByEmail(ctx, email)
}

func (s *userService) Create(ctx context.Context, payload *model.User) (data *model.UserResponse, err error) {
	payload.Password, err = helper.GenerateHashedPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	return s.UserRepository.Create(ctx, payload)
}

func (s *userService) Update(ctx context.Context, payload *model.User, id uuid.UUID) (data *model.UserResponse, err error) {
	_, err = s.UserRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	payload.Password = string(hashedPassword)
	return s.UserRepository.Update(ctx, payload, id)
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.UserRepository.Delete(ctx, id)
}
