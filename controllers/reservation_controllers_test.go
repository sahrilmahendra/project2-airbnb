package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project2/config"
	"project2/constants"
	"project2/middlewares"
	"project2/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type GetReservResponse struct {
	Message string
	Data    models.GetReserv
}

type ResponSuccess struct {
	Status  string
	Message string
	Data    []models.ReservationRequest
}

var (
	mock_data_credit = models.CreditCard{
		ReservationID: 1,
		Typ:           "visa",
		Name:          "sahril",
		Number:        "1100",
		Cvv:           1111,
		Month:         1,
		Year:          21,
	}
	mock_data_reser = models.Reservation{
		UsersID:     1,
		HomestayID:  1,
		Start_date:  "2021-12-1",
		End_date:    "2021-12-2",
		Total_harga: 100000,
	}
)

var (
	mock_data_reservation = models.ReservationRequest{
		Reservation: mock_data_reser,
		Credit:      mock_data_credit,
	}
)

func InsertReser() error {
	if err := config.DB.Save(&mock_data_reservation).Error; err != nil {
		return err
	}
	return nil
}

func TestGetReservationControllersBadRequest(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Bad Request",
		path: "jwt/reservation",
		code: 400,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	middleware.JWT([]byte(constants.SECRET_JWT))(GetReservationControllersTesting())(context)

	var Reserve GetReservResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &Reserve)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/reservation", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, Reserve.Message)
	})

}
func TestGetReservationControllersSuccess(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Successful Operation",
		path: "jwt/reservation",
		code: 200,
	}

	e := InitEcho()
	InsertReser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	middleware.JWT([]byte(constants.SECRET_JWT))(GetReservationControllersTesting())(context)

	var Reserve GetReservResponse
	var Reservet ResponSuccess
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &Reserve)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/reservation", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, Reservet.Message)
		assert.Equal(t, testCases.name, Reserve.Data)
		assert.Equal(t, "sahril", Reservet.Data[0].Credit.Name)
	})

}
