package users

import "context"
import "fmt"

import "github.com/jackc/pgx/v5"

import "github.com/muse/settings"

var createUsersTable string = `
CREATE TABLE users (
	id uuid PRIMARY KEY,
	username varchar(255),
	password varchar(255),
);
`

func GetPsqlConnection(settings settings.PostgresSettings) pgx.Conn {
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

func CreateUsersTable(conn pgx.Conn) {
	// todo: logging
	var result string
	err := conn.QueryRow(createUsersTable).Scan(&result)
	if err != nil {
		fmt.Println("Woops, error during creating users table.")
	}
	fmt.Println(result)
}
