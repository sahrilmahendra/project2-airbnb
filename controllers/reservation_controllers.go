package controllers

import (
	"fmt"
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
	if er != nil || data == "Not Available" || data == 0 {
		fmt.Println("masuk yang ini")
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	} else {

		// mencari berapa hari reservasi
		day := databases.SearchDay(start, end)
		price, id_user_homestay, _ := databases.GetPriceIDuserHomestay(int(Reservation.Reservation.HomestayID), day)

		//cek iduser di homestay
		if id == int(id_user_homestay) {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
		}

		Reservation.Reservation.Total_harga = price
		_, err := databases.CreateReservation(&Reservation)
		if err != nil {
			fmt.Println("masuk yang ini 3")
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
		}

		return c.JSON(http.StatusBadRequest, response.SuccessResponseNonData())
	}
}

func CekReservationControllers(c echo.Context) error {
	cek := models.CekStatus{}
	c.Bind(&cek)

	_, err := databases.GetHomestayById(int(cek.HomestayID))
	data, er := databases.CekStatusReservation(cek.HomestayID, cek.Start_date, cek.End_date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	if er != nil || data == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.AvailableResponse(data))

}

func GetReservationControllers(c echo.Context) error {
	id := middlewares.ExtractTokenId(c)

	data, e := databases.GetReservation(id)
	if e != nil || data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(data))
}

//untuk kebutuhan testing get reservation
func GetReservationControllersTesting() echo.HandlerFunc {
	return GetReservationControllers
}
