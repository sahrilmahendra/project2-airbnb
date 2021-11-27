package databases

import (
	"project2/config"
	"project2/models"
)

// var get_facility_by_id models.GetFacility
var get_facilities []models.GetFacility

// function database untuk membuat data facility baru
func CreateFacility(facility *models.Facility) (interface{}, error) {
	if err := config.DB.Create(&facility).Error; err != nil {
		return nil, err
	}
	return facility, nil
}

// function database untuk menampilkan seluruh data homestay
func GetAllFacilities() (interface{}, error) {
	query := config.DB.Table("facilities").Select("*").Find(&get_facilities)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_facilities, nil
}
