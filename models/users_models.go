package models

import "gorm.io/gorm"

// struktur data users
type Users struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string
}

type GetUser struct {
	ID    int
	Nama  string
	Email string
}
