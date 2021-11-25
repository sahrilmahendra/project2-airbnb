package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	response "project2/responses"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidatorHomestay struct {
	Name_Homestay string `validate:"required"`
	Price         int    `validate:"required,gt=0"`
	Description   string `validate:"required"`
	Address       string `validate:"required"`
}

// controller untuk menambahkan homestay baru
func CreateHomestayControllers(c echo.Context) error {
	new_homestay := models.Homestay{}
	c.Bind(&new_homestay)
	v := validator.New()
	validasi_homestay := ValidatorHomestay{
		Name_Homestay: new_homestay.Name_Homestay,
		Price:         new_homestay.Price,
		Description:   new_homestay.Description,
		Address:       new_homestay.Address,
	}
	err := v.Struct(validasi_homestay)
	if err == nil {
		logged := middlewares.ExtractTokenId(c)
		new_homestay.UsersID = uint(logged)
		_, err = databases.CreateHomestay(&new_homestay)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk menampilkan seluruh data homestay
func GetAllHomestayControllers(c echo.Context) error {
	homestay, err := databases.GetAllHomestay()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(homestay))
}
