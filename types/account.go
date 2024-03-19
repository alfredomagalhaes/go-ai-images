package types

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int `bun:"id,pk,autoincrement"`
	UserID    uuid.UUID
	UserName  string
	UserEmail string
	CreatedAt time.Time `bun:"default:'now()'"`
}
