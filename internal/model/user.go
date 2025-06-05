package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"password"`
}
