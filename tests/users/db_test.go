package users

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/muse/common"
	"github.com/muse/users"
)

func TestGetPsqlConnection(t *testing.T) {
	conn := common.GetPsqlConnection()
	defer conn.Close(context.Background())
	userRepository := users.NewUserRepository(conn, context.Background())

	user := users.NewUser(
		uuid.New(),
		"dummy_username",
		"dummy_password",
	)
	err := userRepository.Create(user)
	if err != nil {
		t.Fatal("Unexpected error")
	}
	userReturned, err := userRepository.Get(user.Id)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	if userReturned != user {
		t.Fatalf("Incorrect result. \n Expected: %s \n Got: %s", user, userReturned)
	}
	users.CleanUsersTable(conn)
}
