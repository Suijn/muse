package common

type ClientError struct {
	Detail string
}

func (o ClientError) Error() string {
	return "Client error"
}

type UnrecoverableError struct {
	Detail string
}

func (o UnrecoverableError) Error() string {
	return "Unrecoverable error"
}
