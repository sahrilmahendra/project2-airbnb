package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	// UsersID    uint
	Start_date string `json:"start_date" form:"start_date"`
	End_date   string `json:"end_date" form:"end_date"`
	Credit     Credit
}

type Credit struct {
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
	Credit      Credit      `json:"credit_card" `
}
