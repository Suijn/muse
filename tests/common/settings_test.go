package common

import (
	"os"
	"testing"

	"github.com/muse/common/settings"
)

func TestNewSettings(t *testing.T) {
	os.Setenv("POSTGRES_HOST", "postgres_host")
	os.Setenv("POSTGRES_PASSWORD", "postgres_password")
	os.Setenv("POSTGRES_USER", "postgres_user")
	os.Setenv("POSTGRES_DB", "postgres_db")
	os.Setenv("POSTGRES_PORT", "5432")

	result := settings.NewSettings()
	expectedPostgresSettings := settings.PostgresSettings{
		POSTGRES_HOST:     "postgres_host",
		POSTGRES_PASSWORD: "postgres_password",
		POSTGRES_USER:     "postgres_user",
		POSTGRES_DB:       "postgres_db",
		POSTGRES_PORT:     5432,
	}
	expected_result := settings.Settings{
		PostgresSettings: expectedPostgresSettings,
	}

	if result != expected_result {
		t.Fatalf("Incorrect result. \n Expected: %s, \n Got: %s", expected_result, result)
	}
}
