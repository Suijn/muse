package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var cleanUsersTableQuery string = `
DELETE FROM users;		
`

func CleanUsersTable(conn *pgx.Conn) {
	res, err := conn.Exec(context.Background(), cleanUsersTableQuery)
	if err != nil {
		fmt.Println("Woops, error during cleaning users table.")
	}
	fmt.Println(res)
}
