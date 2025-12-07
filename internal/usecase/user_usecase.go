package usecase

import "clean-architecture/internal/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetByID(id int) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) GetAll() ([]*domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUsecase) Create(user *domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return u.userRepo.Create(user)
}

func (u *userUsecase) Update(id int, user *domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	user.ID = id
	return u.userRepo.Update(user)
}

func (u *userUsecase) Delete(id int) error {
	return u.userRepo.Delete(id)
}
