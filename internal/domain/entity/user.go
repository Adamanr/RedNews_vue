package entity

type User struct {
	UserId   int64  `json:"usersId"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
