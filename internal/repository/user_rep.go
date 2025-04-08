package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DrusGalkin/forum-auth-grpc/internal/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) entity.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	query := `SELECT id, name, email, password, active FROM users ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Ошибка запроса к пользователям: %w", err)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Active)
		if err != nil {
			return nil, fmt.Errorf("Ошибка получения пользователя: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка итерации по сторокам: %w", err)
	}
	return users, nil
}

func (r *userRepository) GetByID(id int) (entity.User, error) {
	query := `SELECT id, name, email, password, active FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, entity.ErrorUserNotFound
		}
		return entity.User{}, fmt.Errorf("Ошибка получения пользователя по ID: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	query := `SELECT id, name, email, password, active FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, entity.ErrorUserNotFound
		}
		return entity.User{}, fmt.Errorf("Ошибка получения пользователя по Email: %w", err)
	}

	return user, nil
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	query :=
		`INSERT INTO users (name, email, password, active) 
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, name, email, password, active`

	var createdUser entity.User
	err := r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Active,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.Password,
		&createdUser.Active,
	)

	if err != nil {
		return entity.User{}, fmt.Errorf("Ошибка создания пользователя: %w", err)
	}

	return createdUser, nil
}

func (r *userRepository) Update(id int, user entity.User) (entity.User, error) {
	query :=
		`UPDATE users
		 SET name = $1, email = $2, password = $3
		 WHERE id = $4
		 RETURNING id, name, email, password`

	var updateUser entity.User
	err := r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		id,
	).Scan(
		&updateUser.ID,
		&updateUser.Name,
		&updateUser.Email,
		&updateUser.Password,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, entity.ErrorUserNotFound
		}
		return entity.User{}, fmt.Errorf("Ошибка изменения пользователя: %w", err)
	}

	if err = updateUser.HashPassword(); err != nil {
		return entity.User{}, err
	}

	return updateUser, nil
}

func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Ошибка удаления пользователя: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Ошибка получения измененных строк: %w", err)
	}

	if rowsAffected == 0 {
		return entity.ErrorUserNotFound
	}
	return nil
}

func (r *userRepository) CheckPassword(id int, password string) bool {
	user, err := r.GetByID(id)
	if err != nil {
		return false
	}
	return user.VerifyPassword(password)
}
