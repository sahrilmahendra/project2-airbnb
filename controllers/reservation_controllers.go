package controllers

import (
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

	d_user, _ := databases.GetUserById(id)

	v, _ := databases.GetHomestayById(int(homestay_id))
	// cek status reservasi
	data, _ := databases.CekStatusReservation(homestay_id, start, end)
	if v == nil || d_user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	} else if data == 1 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Date"))
	} else if data == "Not Available" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Not Available"))
	} else {

		// mencari berapa hari reservasi
		day := databases.SearchDay(start, end)
		price, id_user_homestay, _ := databases.GetPriceIDuserHomestay(int(Reservation.Reservation.HomestayID), day)

		//cek iduser di homestay
		if id == int(id_user_homestay) {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
		}

		Reservation.Reservation.Total_harga = price
		_, err := databases.CreateReservation(&Reservation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
		}
		return c.JSON(http.StatusBadRequest, response.SuccessResponseNonData("Success Operation"))
	}
}

func CekReservationControllers(c echo.Context) error {
	cek := models.CekStatus{}
	c.Bind(&cek)

	v, _ := databases.GetHomestayById(int(cek.HomestayID))
	data, er := databases.CekStatusReservation(cek.HomestayID, cek.Start_date, cek.End_date)

	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	} else if v == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	} else if data == 1 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Date"))
	} else if data == 0 {
		return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", "Available"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))

}

func GetReservationControllers(c echo.Context) error {
	id := middlewares.ExtractTokenId(c)

	data, er := databases.GetReservation(id)
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	} else if data == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

//untuk kebutuhan testing get reservation
func GetReservationControllersTesting() echo.HandlerFunc {
	return GetReservationControllers
}
