package controllers

import (
	"net/http"
	"project2/helper"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	response "project2/responses"
	"strconv"

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
		lat, lng, _ := helper.GetGeocodeLocations(new_homestay.Address)
		new_homestay.Latitude = lat
		new_homestay.Longitude = lng
		new_homestay.Status = "Available"
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

// controller untuk menampilkan data homestay by id
func GetHomestayByIdControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	product, e := databases.GetHomestayById(id)
	if e != nil || product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(product))
}

// controller untuk memperbarui homestay by id
func UpdateHomestayControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_homestay, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenId(c) // check token
	if logged != int(id_user_homestay) {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	update_homestay := models.Homestay{}
	c.Bind(&update_homestay)
	v := validator.New()
	validasi_homestay := ValidatorHomestay{
		Name_Homestay: update_homestay.Name_Homestay,
		Price:         update_homestay.Price,
		Description:   update_homestay.Description,
		Address:       update_homestay.Address,
	}
	e := v.Struct(validasi_homestay)
	if e == nil {
		logged := middlewares.ExtractTokenId(c)
		update_homestay.UsersID = uint(logged)
		lat, lng, _ := helper.GetGeocodeLocations(update_homestay.Address)
		update_homestay.Latitude = lat
		update_homestay.Longitude = lng
		update_homestay.Status = "Available"
		_, e = databases.UpdateHomestay(id, &update_homestay)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func DeleteHomestayControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_product, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenId(c)
	if uint(logged) != id_user_product {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteHomestay(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
