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
		_, err = databases.CreateHomestay(&new_homestay)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menampilkan seluruh data homestay
func GetAllHomestayControllers(c echo.Context) error {
	homestay, err := databases.GetAllHomestay()
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", homestay))
}

// controller untuk menampilkan data homestay by id
func GetHomestayByIdControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	homestay, e := databases.GetHomestayById(id)
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", homestay))
}

// controller untuk memperbarui homestay by id
func UpdateHomestayControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	homestay, _ := databases.GetHomestayById(id)
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	id_user_homestay, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenId(c) // check token
	if logged != int(id_user_homestay) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
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
		_, e = databases.UpdateHomestay(id, &update_homestay)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

func DeleteHomestayControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	homestay, _ := databases.GetHomestayById(id)
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	id_user_homestay, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenId(c)
	if uint(logged) != id_user_homestay {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	databases.DeleteHomestay(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
