package models

type User struct {
	Email  string   `json:"email"`
	Avatar string   `json:"avatar"`
	Cart   []string `json:"cart"`
}
