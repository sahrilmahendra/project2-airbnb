package databases

import (
	"log"
	"project2/config"
	"project2/models"
)

func CreateReservation(Reser *models.ReservationRequest) (interface{}, error) {

	if err := config.DB.Create(&Reser.Reservation).Error; err != nil {
		return nil, err
	}
	req_reservation := models.Reservation{
		HomestayID: Reser.Reservation.HomestayID,
		Start_date: Reser.Reservation.Start_date,
		End_date:   Reser.Reservation.End_date,
	}

	req_credit := models.CreditCard{
		Typ:    Reser.Credit.Typ,
		Name:   Reser.Credit.Name,
		Number: Reser.Credit.Number,
		Cvv:    Reser.Credit.Cvv,
		Month:  Reser.Credit.Month,
		Year:   Reser.Credit.Year,
	}

	Create_Res := models.ReservationRequest{
		Reservation: req_reservation,
		Credit:      req_credit,
	}
	Reser.Credit.ReservationID = Reser.Reservation.ID
	config.DB.Create(&Reser.Credit)
	return Create_Res, nil
}

func GetPriceIDuserHomestay(id, day int) (int, uint, error) {
	homestay := models.Homestay{}
	err := config.DB.Find(&homestay, id)
	if err.Error != nil {
		return 0, 0, err.Error
	}
	log.Println("harga", homestay.Price)
	return homestay.Price * day, homestay.UsersID, nil
}

// func GetIDUserReservation(id int) (uint, error) {
// 	var reservation models.Reservation
// 	err := config.DB.Find(&reservation, id)
// 	if err.Error != nil {
// 		return 0, err.Error
// 	}
// 	return reservation.UsersID, nil
// }

func GetReservation(id int) (interface{}, error) {
	var get_reservation []models.GetReserv

	query := config.DB.Table("reservations").Select("*").Where("reservations.users_id = ?", id).Find(&get_reservation)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	query_homestay := config.DB.Table("homestays").Select("reservations.users_id,reservations.homestay_id,homestays.name_homestay,reservations.start_date,reservations.end_date,homestays.price,reservations.total_harga").Joins("join reservations on homestays.id = reservations.homestay_id").Find(&get_reservation)
	if query_homestay.Error != nil {
		return nil, query_homestay.Error
	}
	log.Println("gethomestay :", get_reservation[0].Start_date)
	return get_reservation, nil

}
