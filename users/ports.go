package users

type UserRepositoryInterface interface {
	Create(user User) error
}
