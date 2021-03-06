package models

import "gorm.io/gorm"

// struktur data users
type Users struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Token       string
	Homestay    []Homestay
	Reservation []Reservation
}

type GetUser struct {
	ID    uint
	Name  string
	Email string
}

type GetLoginUser struct {
	ID    uint
	Name  string
	Token string
}
