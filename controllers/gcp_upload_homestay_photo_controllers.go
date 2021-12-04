package controllers

import (
	"net/http"
	response "project2/responses"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var storageClient *storage.Client

func UploadFileToGCSBucket(c echo.Context) error {
	bucket := "sahril-bucket"

	var err error

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Can't Connect to Server"))
	}

	f, uploaded_file, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Can't Connect to Server"))
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)
}
