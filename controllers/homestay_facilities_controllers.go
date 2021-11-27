package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	response "project2/responses"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidatorHomestayFacility struct {
	HomestayID uint `validate:"required"`
	FacilityID uint `validate:"required"`
}

// controller untuk menambahkan homestay facility baru
func CreateHomestayFacilityControllers(c echo.Context) error {
	new_homestay_facility := models.Homestay_Facility{}
	c.Bind(&new_homestay_facility)
	v := validator.New()
	validasi_homestay_facility := ValidatorHomestayFacility{
		HomestayID: new_homestay_facility.HomestayID,
		FacilityID: new_homestay_facility.FacilityID,
	}
	err := v.Struct(validasi_homestay_facility)
	var homestay_facility interface{}
	if err == nil {
		logged := middlewares.ExtractTokenId(c)
		id_user_homestay, _ := databases.GetIDUserHomestay(logged)
		if logged == int(id_user_homestay) {
			homestay_facility, err = databases.CreateHomestayFacility(&new_homestay_facility)
		}
	}
	if err != nil || homestay_facility == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk memperbarui homestay facility by id
func UpdateHomestayFacilityControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	update_homestay_facility := models.Homestay_Facility{}
	c.Bind(&update_homestay_facility)
	v := validator.New()
	validasi_homestay_facility := ValidatorHomestayFacility{
		HomestayID: update_homestay_facility.HomestayID,
		FacilityID: update_homestay_facility.FacilityID,
	}
	e := v.Struct(validasi_homestay_facility)
	var homestay_facility interface{}
	if e == nil {
		logged := middlewares.ExtractTokenId(c)
		id_user_homestay, _ := databases.GetIDUserHomestay(logged)
		if logged == int(id_user_homestay) {
			homestay_facility, e = databases.UpdateHomestayFacility(id, &update_homestay_facility)
		}
	}
	if e != nil || homestay_facility == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk mengapus data homestay facility by id
func DeleteHomestayFacilityControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_homestay, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenId(c)
	if uint(logged) != id_user_homestay {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteHomestayFacility(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
