package requests

// Get Token
type Login struct {
	UserNameOrEmail string
	Password        string
}

type Register struct {
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	FullName    string
	Role        int32
	Address     *string
}
