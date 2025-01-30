package main

import "testing"
import "github.com/google/uuid"

func TestUserNew(t *testing.T) {
	id := uuid.New()
	username := "dummy_username"
	password := "dummy_password"

	var user = NewUser(id, username, password)
	expectedUser := User{
		id:       id,
		username: username,
		password: password,
	}
	if expectedUser != user {
		t.Fatalf("Incorrect result, got %s, expected %s", user, expectedUser)
	}
}
