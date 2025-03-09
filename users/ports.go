package users

import (
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	Create(user User) error
	Get(id uuid.UUID) (User, error)
}

type TransactionManager interface {
	Commit() error
}
