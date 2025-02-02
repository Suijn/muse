package users

import "github.com/google/uuid"

type IdFactoryInterface interface {
	getUuid() uuid.UUID
}

/*
Default UUID factory.
Uses "github.com/google/uuid" to generate a random UUID.
*/
type DefaultIdFactory struct{}

func (o DefaultIdFactory) getUuid() uuid.UUID {
	return uuid.New()
}

/*
"Generates" a deterministic UUID.

Use it for testing purposes when you need to generate
a deterministic UUID.

Arguments:
-id: A string. Must be UUID formatted.
*/
type DeterministicIdFactory struct {
	Id string
}

func (o DeterministicIdFactory) getUuid() uuid.UUID {
	return uuid.MustParse(o.Id)
}
