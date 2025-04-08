package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	GetByEmail(email string) (User, error)
	Create(user User) (User, error)
	Update(id int, user User) (User, error)
	Delete(id int) error
	CheckPassword(id int, password string) bool
}

var (
	ErrorUserNotFound  = errors.New("Пользователь не найден")
	ErrorWrongPassword = errors.New("Неперный пароль")
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
	Active   bool   `json:"active"`
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
