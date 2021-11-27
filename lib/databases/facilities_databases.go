package databases

import (
	"log"
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

// function database untuk menampilkan seluruh data facility
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

// function database untuk memperbarui data facility by id
func UpdateFacility(id int, update_facility *models.Facility) (interface{}, error) {
	var facility models.Facility
	query_select := config.DB.Find(&facility, id)
	log.Println("query_sel", query_select.RowsAffected)
	if query_select.Error != nil || query_select.RowsAffected == 0 {
		return 0, query_select.Error
	}
	query_update := config.DB.Model(&facility).Updates(update_facility)
	if query_update.Error != nil {
		return nil, query_update.Error
	}
	return facility, nil
}
