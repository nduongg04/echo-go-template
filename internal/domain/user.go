package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password  string    `json:"-" gorm:"not null" validate:"required,min=6"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type UserService interface {
	Register(user *User) error
	Login(email, password string) (string, error)
	GetProfile(id uint) (*User, error)
	UpdateProfile(user *User) error
}
