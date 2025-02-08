package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var insertUserQuery string = `
	INSERT INTO users (id, username, password)
	VALUES (
		$1,
		$2,
		$3
	);
`

var getUserQuery string = `	
	SELECT id, username, password
	FROM users
	WHERE id = $1;
`

type UserRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewUserRepository(conn *pgx.Conn, ctx context.Context) UserRepository {
	return UserRepository{conn: conn, ctx: ctx}
}

func (o UserRepository) Create(user User) error {
	result, err := o.conn.Exec(
		o.ctx,
		insertUserQuery,
		user.Id.String(),
		user.Username,
		user.Password,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return nil
}

func (o UserRepository) Get(id uuid.UUID) (User, error) {
	var userId string
	var username string
	var password string

	err := o.conn.QueryRow(
		o.ctx,
		getUserQuery,
		id.String(),
	).Scan(&userId, &username, &password)

	if err != nil {
		fmt.Println("Error when getting User")
		return User{}, UserNotFound{}
	}
	return User{
		Id:       uuid.MustParse(userId),
		Username: username,
		Password: password,
	}, nil
}
