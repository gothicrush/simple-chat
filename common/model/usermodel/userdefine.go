package usermodel

type User struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}