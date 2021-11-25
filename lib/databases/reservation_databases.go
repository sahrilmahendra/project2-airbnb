package databases

import (
	"project2/config"
	"project2/models"
)

type Reservation struct {
	Start_date string
	End_date   string
}

type Credit struct {
	Typ    string
	Name   string
	Number string
	Cvv    int
	Month  int
	Year   int
}

type ReservationRequest struct {
	Reservation Reservation
	Credit      Credit
}

func CreateReservation(Reser *models.ReservationRequest) (interface{}, error) {
	if err := config.DB.Create(&Reser.Reservation).Error; err != nil {
		return nil, err
	}
	Reser.Credit.ReservationID = Reser.Reservation.ID
	config.DB.Create(&Reser.Credit)
	return ReservationRequest{
		Reservation{
			Reser.Reservation.Start_date,
			Reser.Reservation.End_date,
		},
		Credit{
			Reser.Credit.Typ,
			Reser.Credit.Name,
			Reser.Credit.Number,
			Reser.Credit.Cvv,
			Reser.Credit.Month,
			Reser.Credit.Year,
		},
	}, nil
}
