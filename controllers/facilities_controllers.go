package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/models"
	response "project2/responses"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan facility baru
func CreateFacilityControllers(c echo.Context) error {
	new_facility := models.Facility{}
	c.Bind(&new_facility)
	v := validator.New()
	err := v.Var(new_facility.Name_Facility, "required")
	if err == nil {
		_, err = databases.CreateFacility(&new_facility)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menampilkan seluruh data facility
func GetAllFacilityControllers(c echo.Context) error {
	facilities, err := databases.GetAllFacilities()
	if facilities == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", facilities))
}

// controller untuk menampilkan data facility by id
func GetFacilityByIdControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	facility, e := databases.GetFacilityById(id)
	if facility == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", facility))
}

// controller untuk memperbarui facility by id
func UpdateFacilityControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	facility, _ := databases.GetFacilityById(id)
	if facility == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	update_facility := models.Facility{}
	c.Bind(&update_facility)
	v := validator.New()
	e := v.Var(update_facility.Name_Facility, "required")
	if e == nil {
		_, e = databases.UpdateFacility(id, &update_facility)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menghapus facility by id
func DeleteFacilityControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	facility, _ := databases.GetFacilityById(id)
	if facility == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	_, e := databases.DeleteFacility(id)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
