package common

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/muse/common/settings"
)

func GetPsqlConnection() *pgx.Conn {
	settings := settings.AppSettings.PostgresSettings
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		settings.POSTGRES_USER,
		settings.POSTGRES_PASSWORD,
		settings.POSTGRES_HOST,
		settings.POSTGRES_PORT,
		settings.POSTGRES_DB,
	)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic("Error when connecting to Postgres.")
	}
	return conn
}
