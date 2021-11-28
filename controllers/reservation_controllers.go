package controllers

import (
	"log"
	"net/http"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	response "project2/responses"

	"github.com/labstack/echo/v4"
)

func CreateReservationControllers(c echo.Context) error {
	Reservation := models.ReservationRequest{}

	c.Bind(&Reservation)
	id := middlewares.ExtractTokenId(c)
	Reservation.Reservation.UsersID = uint(id)

	homestay_id := Reservation.Reservation.HomestayID
	start := Reservation.Reservation.Start_date
	end := Reservation.Reservation.End_date

	// cek status reservasi
	data, er := databases.CekStatusReservation(homestay_id, start, end)
	if er != nil || data == "not available" || data == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	} else {

		// mencari berapa hari reservasi
		day := databases.SearchDay(start, end)
		price, id_user_homestay, _ := databases.GetPriceIDuserHomestay(int(Reservation.Reservation.HomestayID), day)

		log.Println("berapa hari :", day)
		log.Println("id user homesaty", id_user_homestay)

		//cek iduser di homestay
		if id == int(id_user_homestay) {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
		}

		Reservation.Reservation.Status = "not available"
		Reservation.Reservation.Total_harga = price
		_, err := databases.CreateReservation(&Reservation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
		}

		return c.JSON(http.StatusBadRequest, response.SuccessResponseNonData())
	}
}

func CekReservationControllers(c echo.Context) error {
	cek := models.CekStatus{}
	c.Bind(&cek)

	data, er := databases.CekStatusReservation(cek.HomestayID, cek.Start_date, cek.End_date)
	if er != nil || data == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(data))
}

func GetReservationControllers(c echo.Context) error {
	id := middlewares.ExtractTokenId(c)
	log.Println("id  :", id)
	data, e := databases.GetReservation(id)
	if e != nil || data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(data))
}
