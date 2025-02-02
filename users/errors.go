package users

type UserAlreadyExists struct{}

func (o UserAlreadyExists) Error() string {
	return "User already exists"
}
