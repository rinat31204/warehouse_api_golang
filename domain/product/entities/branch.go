package entities

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	id         uuid.UUID
	name       string
	userId     uuid.UUID
	dateCreate time.Time
}
