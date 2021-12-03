package databases

import (
	"project2/config"
	"project2/models"
)

var get_homestay []models.GetHomestay

// function database untuk membuat data homestay baru
func CreateHomestay(homestay *models.Homestay) (interface{}, error) {
	if err := config.DB.Create(&homestay).Error; err != nil {
		return nil, err
	}
	return homestay, nil
}

// function database untuk menampilkan seluruh data homestay
func GetAllHomestay() (interface{}, error) {
	query := config.DB.Table("homestays").Select("*").Where("homestays.deleted_at IS NULL").Find(&get_homestay)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_homestay, nil
}

// function database untuk menampilkan data homestay by id
func GetHomestayById(id int) (interface{}, error) {
	var get_homestay_by_id_user models.GetHomestay
	query := config.DB.Table("homestays").Select("*").Where("homestays.deleted_at IS NULL AND homestays.id = ?", id).Find(&get_homestay_by_id_user)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_homestay_by_id_user, nil
}

// function database untuk memperbarui data homestay by id
func UpdateHomestay(id int, update_homestay *models.Homestay) (interface{}, error) {
	var homestay models.Homestay
	query_select := config.DB.Find(&homestay, id)
	if query_select.Error != nil || query_select.RowsAffected == 0 {
		return 0, query_select.Error
	}
	query_update := config.DB.Model(&homestay).Updates(update_homestay)
	if query_update.Error != nil {
		return nil, query_update.Error
	}
	return homestay, nil
}

// function bantuan untuk mendapatkan id user pada tabel homestay
func GetIDUserHomestay(id int) (uint, error) {
	var homestay models.Homestay
	err := config.DB.Find(&homestay, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return homestay.UsersID, nil
}

// function database untuk menghapus homestay by id
func DeleteHomestay(id int) (interface{}, error) {
	var homestay models.Homestay
	check_homestay := config.DB.Find(&homestay, id).RowsAffected

	err := config.DB.Delete(&homestay).Error
	if err != nil || check_homestay > 0 {
		return nil, err
	}
	return homestay.UsersID, nil
}

// function database untuk menampilkan homestay by id user
func GetHomestayByIdUser(id int) (interface{}, error) {
	var get_homestay_by_id_user []models.GetHomestay
	query := config.DB.Table("homestays").Select("*").Where("homestays.deleted_at IS NULL AND homestays.users_id = ?", id).Find(&get_homestay_by_id_user)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_homestay_by_id_user, nil
}
