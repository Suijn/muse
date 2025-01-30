package main

import "fmt"
import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
}

func NewUser(id uuid.UUID, username string, password string) User {
	// Constructor for User type.
	return User{id, username, password}
}

func (u User) String() string {
	return fmt.Sprintf("User {id: %s, username: %s}", u.id, u.username)
}
