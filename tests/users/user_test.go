package users

import "testing"
import "github.com/google/uuid"

import "github.com/muse/users"

func TestUserNew(t *testing.T) {
	id := uuid.New()
	username := "dummy_username"
	password := "dummy_password"

	var user = users.NewUser(id, username, password)
	expectedUser := users.User{
		Id:       id,
		Username: username,
		Password: password,
	}
	if expectedUser != user {
		t.Fatalf("Incorrect result, got %s, expected %s", user, expectedUser)
	}
}
