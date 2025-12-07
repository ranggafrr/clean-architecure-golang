package repository

import (
	"clean-architecture/internal/domain"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAll() ([]*domain.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users ORDER BY id DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return nil
}

func (r *userRepository) Update(user *domain.User) error {
	query := "UPDATE users SET name = ?, email = ?, updated_at = NOW() WHERE id = ?"
	result, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
