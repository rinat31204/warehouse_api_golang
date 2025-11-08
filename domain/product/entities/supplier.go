package entities

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	id        uuid.UUID
	name      string
	createdAt time.Time
}
