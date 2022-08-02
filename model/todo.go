package entityModel

import "gorm.io/gorm"

type NewTodo struct {
	gorm.Model
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	gorm.Model
	Text string `json:"text"`
	Done bool   `json:"done"`
	//	User *User  `json:"user"`
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}
