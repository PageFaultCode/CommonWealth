// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

import (
	"github.com/google/uuid"
)

// The response to the request
type WebResponse struct {
	ID       uuid.UUID `json:"ID"`
	HttpCode int       `json:"code"`
}
