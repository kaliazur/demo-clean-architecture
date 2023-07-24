package domain

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Age       int    `json:"age"`
	IsAdmin   bool   `json:"is_admin"`
}
