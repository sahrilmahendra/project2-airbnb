package databases

import (
	"project2/config"
	"project2/models"
)

// var get_facility_by_id models.GetFacility
// var get_facilities []models.GetFacility

// function database untuk membuat data facility baru
func CreateFacility(facility *models.Facility) (interface{}, error) {
	if err := config.DB.Create(&facility).Error; err != nil {
		return nil, err
	}
	return facility, nil
}
