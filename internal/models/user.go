package models

// User model info
// @Description User information
// @Description with user id, name, surname and birthday
type User struct {
	Id       int         `json:"id"`
	Name     interface{} `json:"name"`
	Surname  interface{} `json:"surname"`
	Birthday interface{} `json:"birthday"`
}
