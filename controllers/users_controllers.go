package controllers

import (
	"log"
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

// struktur data untuk validasi user
type ValidatorUser struct {
	Nama     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// controller untuk menampilkan data user by id
func GetUserControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	log.Println("id", conv_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	user, e := databases.GetUserById(conv_id)
	if e != nil || user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(user))
}

// controller untuk menambahkan user (registrasi)
func CreateUserControllers(c echo.Context) error {
	new_user := models.Users{}
	c.Bind(&new_user)

	v := validator.New()
	validasi_user := ValidatorUser{
		Nama:     new_user.Nama,
		Email:    new_user.Email,
		Password: new_user.Password,
	}
	err := v.Struct(validasi_user)
	if err == nil {
		new_user.Password, _ = helper.HashPassword(new_user.Password) // generate plan password menjadi hash
		_, err = databases.CreateUser(&new_user)
	}
	check, _ := databases.GetUserByEmail(new_user.Email)
	if err != nil || check > 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(new_user))
}

// controller untuk menghapus user by id
func DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	logged := middlewares.ExtractTokenId(c) // check token
	if logged != conv_id {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	_, e := databases.DeleteUser(conv_id)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk memperbarui data user by id (update)
func UpdateUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	logged := middlewares.ExtractTokenId(c) // check token
	if logged != id {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	users := models.Users{}
	c.Bind(&users)

	v := validator.New()
	validasi_user := ValidatorUser{
		Nama:     users.Nama,
		Email:    users.Email,
		Password: users.Password,
	}
	e := v.Struct(validasi_user)
	if e == nil {
		users.Password, _ = helper.HashPassword(users.Password) // generate plan password menjadi hash
		_, e = databases.UpdateUser(id, &users)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk login dan generate token (by email dan password)
func LoginUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	plan_pass := user.Password
	log.Println(plan_pass)
	token, e := databases.LoginUser(plan_pass, &user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.LoginFailedResponse())
	}
	return c.JSON(http.StatusOK, response.LoginSuccessResponse(token))
}

// controller untuk kebutuhan testing get user
func GetUserControllersTesting() echo.HandlerFunc {
	return GetUserControllers
}

// controller untuk kebutuhan testing update user
func UpdateUserControllersTesting() echo.HandlerFunc {
	return UpdateUserControllers
}

// controller untuk kebutuhan testing delete user
func DeleteUserControllersTesting() echo.HandlerFunc {
	return DeleteUserControllers
}
