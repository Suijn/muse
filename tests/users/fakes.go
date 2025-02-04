package users

import "github.com/muse/users"

type fakeUserRepository struct {
	db map[string]users.User
}

func (f fakeUserRepository) Create(user users.User) error {
	userId := user.Id.String()
	_, exists := f.db[userId]
	if exists {
		return users.UserAlreadyExists{}
	}

	f.db[userId] = user
	return nil
}
