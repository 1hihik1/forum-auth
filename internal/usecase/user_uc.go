package usecase

import (
	"fmt"
	"github.com/DrusGalkin/forum-auth-grpc/internal/entity"
)

type UserUseCase struct {
	repo entity.UserRepository
}

func NewUserUseCase(repo entity.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetAllUsers() ([]entity.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) GetUserByID(id int) (entity.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUseCase) GetUserByEmail(email string) (entity.User, error) {
	return uc.repo.GetByEmail(email)
}

func (uc *UserUseCase) CreateUser(user entity.User) (entity.User, error) {
	if user.Password == "" {
		return entity.User{}, fmt.Errorf("пароль не может быть пустым")
	}

	if err := user.HashPassword(); err != nil {
		return entity.User{}, err
	}

	user.Active = true

	return uc.repo.Create(user)
}

func (uc *UserUseCase) UpdateUser(id int, user entity.User) (entity.User, error) {
	return uc.repo.Update(id, user)
}

func (uc *UserUseCase) DeleteUser(id int) error {
	return uc.repo.Delete(id)
}

func (uc *UserUseCase) CheckPassword(id int, password string) bool {
	return uc.repo.CheckPassword(id, password)
}
