package entities

import "github.com/google/uuid"

type Product struct {
	id   uuid.UUID
	name string
}
