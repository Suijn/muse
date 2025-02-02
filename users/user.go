package users

import "fmt"
import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Username string
	Password string
}

func NewUser(id uuid.UUID, username string, password string) User {
	// Constructor for User type.
	return User{id, username, password}
}

func (u User) String() string {
	return fmt.Sprintf("User {Id: %s, Username: %s}", u.Id, u.Username)
}
