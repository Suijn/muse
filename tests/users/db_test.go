package users

import "testing"
import "github.com/google/uuid"

import "github.com/muse/users"
import "github.com/muse/settings"

func TestGetPsqlConnection(t *testing.T) {
	conn := users.GetPsqlConnection(settings.AppSettings.PostgresSettings)
	users.CreateUsersTable(conn)
	userRepo := users.NewUserRepository(conn)

	user := users.NewUser(
		uuid.New(),
		"dummy_username",
		"dummy_password",
	)
	userRepo.Create(user)
	// query for user (can be Repo layer)

	// teardown (drop db)
}
