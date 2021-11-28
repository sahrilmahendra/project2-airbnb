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
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk menampilkan seluruh data facility
func GetAllFacilityControllers(c echo.Context) error {
	facilities, err := databases.GetAllFacilities()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(facilities))
}

// controller untuk menampilkan data facility by id
func GetFacilityByIdControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	facility, e := databases.GetFacilityById(id)
	if e != nil || facility == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(facility))
}

// controller untuk memperbarui facility by id
func UpdateFacilityControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	update_facility := models.Facility{}
	c.Bind(&update_facility)
	v := validator.New()
	e := v.Var(update_facility.Name_Facility, "required")
	var facility_rowaffected interface{}
	if e == nil {
		facility_rowaffected, e = databases.UpdateFacility(id, &update_facility)
	}
	if e != nil || facility_rowaffected == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk menghapus facility by id
func DeleteFacilityControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	facility, e := databases.DeleteFacility(id)
	if facility == 0 || e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
