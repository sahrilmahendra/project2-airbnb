package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/models"
	response "project2/responses"

	"github.com/labstack/echo/v4"
)

func CreateReservationControlllers(c echo.Context) error {
	Reservation := models.ReservationRequest{}

	c.Bind(&Reservation)

	reser, er := databases.CreateReservation(&Reservation)
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusBadRequest, response.SuccessResponseData(reser))
}
