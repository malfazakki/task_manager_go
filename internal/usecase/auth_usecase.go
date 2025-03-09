package usecase

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo}
}

func (u *authUsecase) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return u.userRepo.CreateUser(user)
}

func (u *authUsecase) Login(email, password string) (*models.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
