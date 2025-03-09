package users

import (
	"context"

	"github.com/muse/common"
)

func getUserRepository() UserRepositoryInterface {
	conn := common.GetPsqlConnection()
	return NewUserRepository(conn, context.Background())
}

func getIdFactory() IdFactoryInterface {
	return DefaultIdFactory{}
}
