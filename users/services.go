package users

func CreateUser(
	userRepository UserRepositoryInterface,
	idFactory IdFactoryInterface,
	username string,
	password string,
) error {
	user := NewUser(idFactory.getUuid(), username, password)
	err := userRepository.Create(user)
	if err != nil {
		// todo: switch for errors.
		// todo: logging.
		return CreateUserFailed{}
	}
	return nil
}
