package domain

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrNameRequired   = errors.New("name is required")
	ErrEmailRequired  = errors.New("email is required")
	ErrEmailDuplicate = errors.New("email already exists")
)
