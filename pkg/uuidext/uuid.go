package uuidext

import "github.com/google/uuid"

func NewUUID() uuid.UUID {
	return uuid.New()
}

func NewUUIDStr() string {
	return uuid.NewString()
}
