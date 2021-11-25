package databases

import (
	"project2/config"
	"project2/helper"
	"project2/models"
)

// var get_homestay models.GetHomestay
// var get_homestay []models.GetHomestay

// function database untuk membuat data homestay baru
func CreateHomestay(homestay *models.Homestay) (interface{}, error) {
	lat, lng, err := helper.GetGeocodeLocations(homestay.Address)
	if err != nil {
		return nil, err
	}
	homestay.Latitude = lat
	homestay.Longitude = lng
	homestay.Status = "Available"
	if err := config.DB.Create(&homestay).Error; err != nil {
		return nil, err
	}
	return homestay, nil
}
