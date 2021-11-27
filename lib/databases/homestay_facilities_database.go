package databases

import (
	"project2/config"
	"project2/models"
)

// function untuk menambahkan homestay facility baru
func CreateHomestayFacility(homestay_facility *models.Homestay_Facility) (interface{}, error) {
	// facility, _ := GetFacilityById(int(homestay_facility.FacilityID))

	// if facility != nil {
	// 	return nil, nil
	// }
	if e := config.DB.Create(&homestay_facility).Error; e != nil {
		return nil, e
	}
	return homestay_facility, nil
}

// function database untuk memperbarui data homestayfacility by id
func UpdateHomestayFacility(id int, update_homestay_facility *models.Homestay_Facility) (interface{}, error) {
	var homestay_facility models.Homestay_Facility
	query_select := config.DB.Find(&homestay_facility, id)
	if query_select.Error != nil || query_select.RowsAffected == 0 {
		return nil, query_select.Error
	}
	query_update := config.DB.Model(&homestay_facility).Updates(update_homestay_facility)
	if query_update.Error != nil {
		return nil, query_update.Error
	}
	return homestay_facility, nil
}
