// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

import "strings"

// Standard verbs subject/topics
const (
	Read             = "read"
	Create           = "create"
	Update           = "update"
	Delete           = "delete"
	AFM              = "afm"
	Status           = "status"
	Database         = "database"
	Query            = "query"
	SessionDatabase  = "session"
	CompanyIndex     = 0
	ApplicationIndex = 1
	CommandIndex     = 2
)

// NatsTopics are topics that will be handled by each service
type NatsTopics struct {
	Name   string `yaml:"name"`   // the name of the topic to subscribe to (subject)
	Create bool   `yaml:"create"` // the create verb if needed or not ignored if always
	Read   bool   `yaml:"read"`   // the read verb if needed or not ignored if always
	Update bool   `yaml:"update"` // the update verb if needed or not ignored if always
	Delete bool   `yaml:"delete"` // the delete verb if needed or not ignored if always
	Always bool   `yaml:"always"` // the always verb that subscribes to the given name only
}

type StatusData struct {
	ApplicationName    string `json:"app_name"`
	ApplicationHealth  int    `json:"app_health"`
	ApplicationMessage string `json:"app_msg"`
}

type DatabaseQuery struct {
	Query string `json:"query"`
	Error error  `json:"error"`
}

type DatabaseResponse struct {
	Entries map[string]any
}

// CreateSubject create a subject from a number of pieces
func CreateSubject(pieces []string) string {
	return strings.Join(pieces, "/")
}

// StripBrackets removes the starting and ending brackets if any
func StripBrackets(value string) string {
	modified := strings.ReplaceAll(value, "[", "")

	modified = strings.ReplaceAll(modified, "]", "")

	return modified
}
