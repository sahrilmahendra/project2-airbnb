package databases

import (
	"fmt"
	"project2/config"
	"project2/models"
)

// function untuk menambahkan homestay facility baru
func CreateHomestayFacility(homestay_facility *models.Homestay_Facility) (interface{}, error) {
	// facility, _ := GetFacilityById(int(homestay_facility.FacilityID))

	// if facility != nil {
	// 	return nil, nil
	// }
	cek := CekHomestayFacilityById(int(homestay_facility.HomestayID), int(homestay_facility.FacilityID))
	fmt.Println("cek", cek)
	if cek != "Data Is Available" {
		e := config.DB.Create(&homestay_facility).Error
		if e != nil || cek == "Query Error" {
			return nil, e
		}
		return homestay_facility, nil
	} else {
		return cek, nil
	}
}

// function database untuk memperbarui data homestay facility by id
func UpdateHomestayFacility(id int, update_homestay_facility *models.Homestay_Facility) (interface{}, error) {
	var homestay_facility models.Homestay_Facility
	cek := CekHomestayFacilityById(int(update_homestay_facility.HomestayID), int(update_homestay_facility.FacilityID))
	fmt.Println("cek", cek)
	if cek != "Data Is Available" {
		query_select := config.DB.Find(&homestay_facility, id)
		if query_select.Error != nil || query_select.RowsAffected == 0 {
			return nil, query_select.Error
		}
		query_update := config.DB.Model(&homestay_facility).Updates(update_homestay_facility)
		if query_update.Error != nil || cek == "Query Error" {
			return nil, query_update.Error
		}
		return homestay_facility, nil
	} else {
		return cek, nil
	}

}

func GetAllHomestayFacility() (interface{}, error) {
	var homestay_facility []models.Get_Homestay_Facility
	query := config.DB.Table("homestay_facilities").Select("*").Where("deleted_at IS NULL").Find(&homestay_facility)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return homestay_facility, nil
}

// function database untuk menghapus data homestay facility by id
func DeleteHomestayFacility(id int) (interface{}, error) {
	var homestay_facility models.Homestay_Facility
	check_homestay_facility := config.DB.Find(&homestay_facility, id).RowsAffected

	err := config.DB.Delete(&homestay_facility).Error
	if err != nil || check_homestay_facility > 0 {
		return nil, err
	}
	return homestay_facility.ID, nil
}

func CekHomestayFacilityById(id_home, id_facility int) string {
	var homestay_facility models.Homestay_Facility
	err := config.DB.Table("homestay_facilities").Select("*").Where("homestay_facilities.homestay_id = ? && homestay_facilities.facility_id = ? ", id_home, id_facility).Find(&homestay_facility)
	fmt.Println("coba cek isi dhome faci", err)
	if err.RowsAffected <= 0 {
		return "Data Not Found"
	} else if err.Error != nil {
		return "Query Error"
	}
	return "Data Is Available"
}
