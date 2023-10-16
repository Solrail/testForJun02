package models

// User model info
type User struct {
	// Id is the unique
	Id int `json:"id" swaggerignore:"true"`
	// Name is the name
	Name interface{} `json:"name"`
	// Surname is the surname
	Surname interface{} `json:"surname"`
	// Birthday is the date of birthday
	// @format date
	Birthday interface{} `json:"birthday"`
}

type Login struct {
	// Login is the name
	Name interface{} `json:"name"`
	// Password is the name
	Password interface{} `json:"password"`
}
