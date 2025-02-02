package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/muse/users"
)

func TestCreateUser(t *testing.T) {
	username := "dummy_username"
	password := "dummy_password"

	db := make(map[string]users.User)
	userRepo := fakeUserRepository{db}
	idFactory := users.DefaultIdFactory{}
	result := users.CreateUser(userRepo, idFactory, username, password)

	if result != nil {
		t.Fatalf("Incorrect result. Expected: %v, Got: %t", nil, result)
	}

	usersNumber := len(userRepo.db)
	if usersNumber != 1 {
		t.Fatalf("Incorrect result, Expected: %d, Got: %d", 1, usersNumber)
	}
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	id := uuid.New()
	username := "dummy_username"
	password := "dummy_password"
	db := make(map[string]users.User)

	user := users.NewUser(id, username, password)
	db[user.Id.String()] = user

	userRepo := fakeUserRepository{db}
	idFactory := users.DeterministicIdFactory{id.String()}

	result := users.CreateUser(userRepo, idFactory, username, password)
	if result == nil {
		t.Fatalf("Incorrect result. Expected: %t, Got: %t", users.CreateUserFailed{}, result)
	}

	usersNumber := len(userRepo.db)
	if usersNumber != 1 {
		t.Fatalf("Incorrect result, Expected: %d, Got: %d", 1, usersNumber)
	}
}
