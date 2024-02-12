package models

type TokenType string
type Permission string

const (
	USER_SIGN_UP     TokenType = "USER_SIGN_UP"
	EMPLOYEE_SIGN_UP TokenType = "EMPLOYEE_SIGN_UP"
	FORGOT_PASSWORD  TokenType = "FORGOT_PASSWORD"
	EMAIL_CHANGE     TokenType = "EMAIL_CHANGE"
)

const (
	AdminPermission Permission = "admin"
	UserPermission  Permission = "user"
	GuestPermission Permission = "guest"
)

type User struct {
	User_UID    uint       `json:"user_uid"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Permission  Permission `json:"permission"`
	Is_verified bool       `json:"is_verified"`
}

type UserSignUp struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
}

type TokenProfile struct {
	Message  string `json:"message"`
	Validity int    `json:"validity"`
}

type TokenStore struct {
	Email     string    `json:"email"`
	Token     int       `json:"token"`
	TokenType TokenType `json:"token_type"`
}
