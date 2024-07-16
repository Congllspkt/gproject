package services

import "gproject/internal/repos"

type UserService struct {
	userRepo *repos.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repos.NewUserRepo(),
	}
}

func (us *UserService) GetInfouser() string {
	return us.userRepo.GetInfouser()
}