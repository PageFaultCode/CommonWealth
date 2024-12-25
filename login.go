// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

const (
	LoginSubject      = "login"
	LogoutSubject     = "logout"
	UserNameParameter = "user_name"
	TokeParameter     = "token"
	CheckLogin        = "select * from " + UserTableName + " where user_name='%s' and password='%s'"
)

// LoginUser struct to define what is needed to log in a user
type LoginUser struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
