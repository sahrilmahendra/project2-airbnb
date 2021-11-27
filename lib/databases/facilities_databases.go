package databases

import (
	"project2/config"
	"project2/models"
)

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

// function database untuk menampilkan data facility by id
func GetFacilityById(id int) (interface{}, error) {
	get_facility_by_id := models.GetFacility{}
	query := config.DB.Table("facilities").Select("facilities.id, facilities.name_facility").Where("facilities.id = ?", id).Find(&get_facility_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_facility_by_id, nil
}
