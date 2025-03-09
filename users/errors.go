package users

type UserAlreadyExists struct{}

func (o UserAlreadyExists) Error() string {
	return "User already exists"
}

type UserNotFound struct{}

func (o UserNotFound) Error() string {
	return "User not found."
}

type CreateUserFailed struct{}

func (o CreateUserFailed) Error() string {
	return "Create user failed."
}
