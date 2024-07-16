package repos

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfouser() string {
	return "Cong Nguyen"
}