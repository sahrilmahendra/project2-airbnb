package databases

import (
	"project2/config"
	"project2/models"
	"time"
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
	return homestay.Price * day, homestay.UsersID, nil
}

func GetReservation(id int) (interface{}, error) {
	var get_reservation []models.GetReserv

	query_homestay := config.DB.Table("homestays").Select("reservations.users_id,reservations.homestay_id,homestays.name_homestay,reservations.start_date,reservations.end_date,homestays.price,reservations.total_harga").Joins("join reservations on homestays.id = reservations.homestay_id").Where("reservations.users_id = ?", id).Find(&get_reservation)
	if query_homestay.Error != nil || query_homestay.RowsAffected == 0 {
		return nil, query_homestay.Error
	}
	return get_reservation, nil

}

func CekStatusReservation(id_home uint, cek_in, cek_out string) (interface{}, error) {
	var cek []models.Reservation
	var hasil string

	if CekTimeBefore(cek_in, cek_out) == true {
		err := config.DB.Table("reservations").Select("*").Where("reservations.homestay_id = ?", id_home).Find(&cek)
		if err.Error != nil || err.RowsAffected == 0 {
			return 0, err.Error
		}

		for i, _ := range cek {
			hasil = SearchAvailableDay(cek[i].Start_date, cek[i].End_date, cek_in, cek_out)
			if hasil == "Not Available" {
				break
			}
		}
		return hasil, nil
	}
	return 0, nil
}

func SearchAvailableDay(in, out, cek_in, cek_out string) string {
	format := "2006-01-02"
	cek_start, _ := time.Parse(format, cek_in)
	cek_end, _ := time.Parse(format, cek_out)
	start, _ := time.Parse(format, in)
	end, _ := time.Parse(format, out)

	hasil := "Available"
	if (start.Before(cek_start) && end.After(cek_start)) || (start.Before(cek_end) && end.After(cek_end)) {
		hasil = "Not Available"
		return hasil
	} else if start.Equal(cek_start) || end.Equal(cek_start) || start.Equal(cek_end) || end.Equal(cek_end) {
		hasil = "Not Available"
		return hasil
	}
	return hasil
}

func SearchDay(in, out string) int {
	format := "2006-01-02"
	start, _ := time.Parse(format, in)
	end, _ := time.Parse(format, out)

	diff := end.Sub(start)
	return int(diff.Hours() / 24) // number of days
}

func CekTimeBefore(cek_start, cek_end string) bool {
	format := "2006-01-02"
	start, _ := time.Parse(format, cek_start)
	end, _ := time.Parse(format, cek_end)
	if start.Before(end) && time.Now().Before(start) {
		return true
	}
	return false
}
