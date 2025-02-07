package users

import "fmt"
import "context"
import "github.com/jackc/pgx/v5"

var insertUserQuery string = `
	INSERT INTO users (id, username, password)
	VALUES (
		'$1',
		'$2',
		'$3'
	);
`

type UserRepository struct {
	conn pgx.Conn
}

func NewUserRepository(conn pgx.Conn) UserRepository {
	return UserRepository{conn: conn}
}

func (o UserRepository) Create(user User) error {
	var result string
	err := o.conn.QueryRow(
		context.Background(),
		insertUserQuery,
		user.Id.String(),
		user.Username,
		user.Password,
	).Scan(&result)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
