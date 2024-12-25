// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

import (
	"github.com/google/uuid"
)

const (
	SessionTableName = "session"
)

type SessionEntry struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Value  string    `json:"value"`
	Active bool      `json:"active"`
}

type Session struct {
	SessionEntry SessionEntry `db:"session"`
}
