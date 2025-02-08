package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/muse/common/settings"
)

var createUsersTable string = `
CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY,
	username varchar(255),
	password varchar(255)
);
`

var dropUsersTableQuery string = "drop table users"

func GetPsqlConnection(settings settings.PostgresSettings) *pgx.Conn {
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

func CreateUsersTable(conn *pgx.Conn) {
	// todo: logging
	err := conn.Ping(context.Background())
	if err != nil {
		fmt.Println("PING FAIL")
		fmt.Println(err)
	}
	res, err := conn.Exec(context.Background(), createUsersTable)
	if err != nil {
		fmt.Println("Woops, error during creating users table.")
	}
	fmt.Println(res)
}

func DropUsersTable(conn *pgx.Conn) {
	res, err := conn.Exec(context.Background(), dropUsersTableQuery)
	if err != nil {
		fmt.Println("Woops, error during dropping users table.")
	}
	fmt.Println(res)
}
