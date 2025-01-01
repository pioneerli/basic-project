package service

import (
	"basic-project/internal/domain"
	"basic-project/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

var ErrDuplicateEmail = repository.ErrDuplicateEmail

var (
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {

	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {

	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 检查密码对不对
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if !strings.EqualFold(u.Password, password) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}
