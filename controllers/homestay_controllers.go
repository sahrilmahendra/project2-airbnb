package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"project2/helper"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	response "project2/responses"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

type ValidatorHomestay struct {
	Name_Homestay string `validate:"required"`
	Price         int    `validate:"required,gt=0"`
	Description   string `validate:"required"`
	Address       string `validate:"required"`
	File          string `validate:"required"`
	Url           string `validate:"required"`
}

var storageClient *storage.Client

// controller untuk menambahkan homestay baru
func CreateHomestayControllers(c echo.Context) error {
	new_homestay := models.Homestay{}
	c.Bind(&new_homestay)

	bucket := "sahril-bucket"

	var err error

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Can't Connect to Server"))
	}

	f, uploaded_file, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	defer f.Close()

	// cek extension
	// sahril.mahendra.jpg || png || jpeg -> <?php ?> pakai explode
	ext := strings.Split(uploaded_file.Filename, ".")
	extension := ext[len(ext)-1]
	check_extension := strings.ToLower(extension)
	// log.Println("extension", extension)
	if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("File Extention Not Allowed"))
	}

	// check size file
	// log.Println("size file", uploaded_file.Size)
	if uploaded_file.Size == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Illegal File"))

		// 1 megabytes = 1048576 bytes (in binary)
	} else if uploaded_file.Size > 1050000 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Size File Too Big"))
	}
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	name_file := strings.ReplaceAll(new_homestay.Name_Homestay, " ", "-")
	uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", name_file, formatted, extension)
	new_homestay.File = uploaded_file.Filename
	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	new_homestay.Url = fmt.Sprintf("%v", u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	v := validator.New()
	validasi_homestay := ValidatorHomestay{
		Name_Homestay: new_homestay.Name_Homestay,
		Price:         new_homestay.Price,
		Description:   new_homestay.Description,
		Address:       new_homestay.Address,
		File:          new_homestay.File,
		Url:           new_homestay.Url,
	}
	err = v.Struct(validasi_homestay)
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

	bucket := "sahril-bucket"

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Can't Connect to Server"))
	}

	f, uploaded_file, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	defer f.Close()

	// cek extension
	// sahril.mahendra.jpg || png || jpeg -> <?php ?> pakai explode
	ext := strings.Split(uploaded_file.Filename, ".")
	extension := ext[len(ext)-1]
	check_extension := strings.ToLower(extension)
	// log.Println("extension", extension)
	if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("File Extention Not Allowed"))
	}

	// check size file
	// log.Println("size file", uploaded_file.Size)
	if uploaded_file.Size == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Illegal File"))

		// 1 megabytes = 1048576 bytes (in binary)
	} else if uploaded_file.Size > 1050000 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Size File Too Big"))
	}
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	name_file := strings.ReplaceAll(update_homestay.Name_Homestay, " ", "-")
	uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", name_file, formatted, extension)
	update_homestay.File = uploaded_file.Filename
	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	update_homestay.Url = fmt.Sprintf("%v", u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	v := validator.New()
	validasi_homestay := ValidatorHomestay{
		Name_Homestay: update_homestay.Name_Homestay,
		Price:         update_homestay.Price,
		Description:   update_homestay.Description,
		Address:       update_homestay.Address,
		File:          update_homestay.File,
		Url:           update_homestay.Url,
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

func GetMyHomestayControllers(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)
	if logged == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	homestay, e := databases.GetHomestayByIdUser(logged)
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", homestay))
}

func GetHomestayByAddressControllers(c echo.Context) error {
	address := c.Param("id")
	log.Println("address from controller", address)
	homestay, e := databases.GetHomestayByAddress(address)
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", homestay))
}
