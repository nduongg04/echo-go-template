package service

import (
	"errors"
	"time"

	"echo-store-api/internal/domain"
	"echo-store-api/pkg/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(user *domain.User) error {
	// Check if user already exists
	existingUser, err := u.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.userRepo.Create(user)
}

func (u *userService) Login(email, password string) (string, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(utils.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *userService) GetProfile(id uint) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userService) UpdateProfile(user *domain.User) error {
	existingUser, err := u.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}

	// Update only allowed fields
	existingUser.Name = user.Name
	existingUser.UpdatedAt = time.Now()

	return u.userRepo.Update(existingUser)
}
