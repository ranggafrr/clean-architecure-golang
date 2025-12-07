package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

type UserUsecase interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(id int, user *User) error
	Delete(id int) error
}
