package settings

import "os"
import "fmt"
import "strconv"

type Settings struct {
	PostgresSettings PostgresSettings
}

func NewSettings() Settings {
	return Settings{
		PostgresSettings: NewPostgresSettings(),
	}
}

func (o Settings) String() string {
	return fmt.Sprintf("NewSettings {PostgresSettings: %s}", o.PostgresSettings)
}

type PostgresSettings struct {
	POSTGRES_HOST     string
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_DB       string
	POSTGRES_PORT     int64
}

func NewPostgresSettings() PostgresSettings {
	postgresPortStr := os.Getenv("POSTGRES_PORT")
	var postgresPort int64 = 0
	if postgresPortStr != "" {
		var err error
		postgresPort, err = strconv.ParseInt(postgresPortStr, 10, 0)
		if err != nil {
			panic("Error when parsing 'POSTGRES_PORT'")
		}
	}
	return PostgresSettings{
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_PORT:     postgresPort,
	}
}

func (o PostgresSettings) String() string {
	return fmt.Sprintf("PostgresSettings {POSTGRES_HOST: %s, POSTGRES_PASSWORD: %s, POSTGRES_USER: %s, POSTGRES_DB: %s, POSTGRES_PORT: %d}", o.POSTGRES_HOST, o.POSTGRES_PASSWORD, o.POSTGRES_USER, o.POSTGRES_DB, o.POSTGRES_PORT)
}

var AppSettings = NewSettings()
