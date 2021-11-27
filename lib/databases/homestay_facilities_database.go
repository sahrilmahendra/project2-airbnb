package databases

import (
	"project2/config"
	"project2/models"
)

// function untuk menambahkan homestay facility baru
func CreateHomestayFacility(homestay_facility *models.Homestay_Facility) (interface{}, error) {
	if err := config.DB.Create(&homestay_facility).Error; err != nil {
		return nil, err
	}
	return homestay_facility, nil
}
