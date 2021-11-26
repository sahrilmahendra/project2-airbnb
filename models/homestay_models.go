package models

import "gorm.io/gorm"

// struktur data homestay
type Homestay struct {
	gorm.Model
	Name_Homestay string  `json:"name_homestay" form:"name_homestay"`
	Price         int     `json:"price" form:"price"`
	Description   string  `json:"description" form:"description"`
	Address       string  `json:"address" form:"address"`
	Latitude      float64 `json:"latitude" form:"latitude"`
	Longitude     float64 `json:"longitude" form:"longitude"`
	Status        string  `status:"status" form:"status"`
	UsersID       uint
	Reservation   []Reservation
}

type GetHomestay struct {
	ID            uint
	Name_Homestay string
	Price         int
	Description   string
	Address       string
	Latitude      float64
	Longitude     float64
}
