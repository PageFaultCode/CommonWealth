// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

import (
	"time"

	"github.com/google/uuid"
)

// UserSubject so we end up with user/create user/read, etc.
const (
	UserSubject   = "user"
	UserTableName = "user"
)

type UserEntry struct {
	ID             uuid.UUID  `db:"id"`
	FirstName      string     `db:"first_name"`
	LastName       string     `db:"last_name"`
	UserName       string     `db:"user_name"`
	Password       string     `db:"password"`
	PasswordChange *time.Time `db:"password_change"`
	Created        time.Time  `db:"created"`
	Updated        time.Time  `db:"updated"`
	Active         bool       `db:"active"`
}

// User is for storing session information
type User struct {
	ID            uuid.UUID
	LastActivity  time.Time
	Authenticated bool
}
