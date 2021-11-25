package databases

import (
	"project2/config"
	"project2/helper"
	"project2/models"
)

var get_homestay_by_id models.GetHomestay
var get_homestay []models.GetHomestay

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

// function database untuk menampilkan seluruh data homestay
func GetAllHomestay() (interface{}, error) {
	query := config.DB.Table("homestays").Select("*").Find(&get_homestay)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_homestay, nil
}

// function database untuk menampilkan data homestay by id
func GetHomestayById(id int) (interface{}, error) {
	query := config.DB.Table("homestays").Select("*").Where("homestays.id = ?", id).Find(&get_homestay_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_homestay_by_id, nil
}
