package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UsersID     uint
	HomestayID  uint   `json:"homestay_id" form:"homestay_id"`
	Start_date  string `json:"start_date" form:"start_date"`
	End_date    string `json:"end_date" form:"end_date"`
	Total_harga int
	Credit      CreditCard
}

type CreditCard struct {
	ReservationID uint
	Typ           string `json:"typ" form:"typ"`
	Name          string `json:"name" form:"name"`
	Number        string `json:"number" form:"number"`
	Cvv           int    `json:"cvv" form:"cvv"`
	Month         int    `json:"month" form:"month"`
	Year          int    `json:"year" form:"year"`
}

type ReservationRequest struct {
	Reservation Reservation `json:"reservation" `
	Credit      CreditCard  `json:"credit_card" `
}

type GetReserv struct {
	UsersID       uint
	HomestayID    uint
	Name_Homestay string `json:"name_homestay" form:"name_homestay"`
	Start_date    string
	End_date      string
	Price         int `json:"price" form:"price"`
	Total_harga   int
}

type CekStatus struct {
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Start_date string `json:"start_date" form:"start_date"`
	End_date   string `json:"end_date" form:"end_date"`
}
