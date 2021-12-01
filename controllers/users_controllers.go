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
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// controller untuk menampilkan seluruh data users
func GetAllUsersControllers(c echo.Context) error {
	users, err := databases.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", users))
}

// controller untuk menampilkan data user by id
func GetUserControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	log.Println("id", conv_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	user, e := databases.GetUserById(conv_id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", user))
}

// controller untuk menambahkan user (registrasi)
func CreateUserControllers(c echo.Context) error {
	new_user := models.Users{}
	c.Bind(&new_user)

	v := validator.New()
	validasi_user := ValidatorUser{
		Name:     new_user.Name,
		Email:    new_user.Email,
		Password: new_user.Password,
	}
	err := v.Struct(validasi_user)
	if err == nil {
		new_user.Password, _ = helper.HashPassword(new_user.Password) // generate plan password menjadi hash
		_, err = databases.CreateUser(&new_user)
	}
	// check, _ := databases.GetUserByEmail(new_user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menghapus user by id
func DeleteUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}

	user, _ := databases.GetUserById(id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}

	logged := middlewares.ExtractTokenId(c) // check token
	if logged != id {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	databases.DeleteUser(id)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk memperbarui data user by id (update)
func UpdateUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}

	user, _ := databases.GetUserById(id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}

	logged := middlewares.ExtractTokenId(c) // check token
	if logged != id {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	users := models.Users{}
	c.Bind(&users)

	v := validator.New()
	validasi_user := ValidatorUser{
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
	}
	e := v.Struct(validasi_user)
	if e == nil {
		users.Password, _ = helper.HashPassword(users.Password) // generate plan password menjadi hash
		_, e = databases.UpdateUser(id, &users)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk login dan generate token (by email dan password)
func LoginUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	plan_pass := user.Password
	log.Println(plan_pass)
	token, e := databases.LoginUser(plan_pass, &user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Email or Password Incorrect"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Login Success", token))
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
