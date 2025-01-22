// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

const (
	LoginSubject      = "login"
	LogoutSubject     = "logout"
	UserNameParameter = "user_name"
	TokenParameter    = "token"
)

// LoginUser struct to define what is needed to log in a user
type LoginUser struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
