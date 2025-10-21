package repo

type UserRepo struct {
	// Add fields if necessary
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserInfo() string {
	return "User"
}
