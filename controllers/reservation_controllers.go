package controllers

import (
	"log"
	"net/http"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	response "project2/responses"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateReservationControllers(c echo.Context) error {
	Reservation := models.ReservationRequest{}

	c.Bind(&Reservation)
	id := middlewares.ExtractTokenId(c)
	Reservation.Reservation.UsersID = uint(id)

	start := Reservation.Reservation.Start_date
	end := Reservation.Reservation.End_date
	day := SearchDay(start, end)

	price, id_user_homestay, _ := databases.GetPriceIDuserHomestay(int(Reservation.Reservation.HomestayID), day)

	log.Println("berapa hari :", day)
	log.Println("id user homesaty", id_user_homestay)
	//cek iduser di homestay
	if id == int(id_user_homestay) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	Reservation.Reservation.Total_harga = price
	_, er := databases.CreateReservation(&Reservation)
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusBadRequest, response.SuccessResponseNonData())
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

func SearchDay(in, out string) int {
	format := "2006-01-02"
	start, _ := time.Parse(format, in)
	end, _ := time.Parse(format, out)

	diff := end.Sub(start)

	return int(diff.Hours() / 24) // number of days
}
