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
		//// a) Timeouts?
		//// b) Unique constraints -> client error (409 Conflict)

		// todo: logging.
		return CreateUserFailed{}
	}
	return nil
}
