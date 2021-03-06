package databases

import (
	"project2/config"
	"project2/middlewares"
	"project2/models"

	"golang.org/x/crypto/bcrypt"
)

var user models.Users

// function database untuk menampilkan seluruh data users
func GetAllUsers() (interface{}, error) {
	var users []models.GetUser
	query := config.DB.Table("users").Select("*").Find(&users)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return users, nil
}

// function database untuk menampilkan user by id
func GetUserById(id int) (interface{}, error) {
	users := models.Users{}
	var get_user models.GetUser

	err := config.DB.Find(&users, id)
	rows_affected := config.DB.Find(&users, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	get_user.ID = users.ID
	get_user.Name = users.Name
	get_user.Email = users.Email
	return get_user, nil
}

// function database untuk menambahkan user baru (registrasi)
func CreateUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk menghapus user by id
func DeleteUser(id int) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// function database untuk memperbarui data user by id
func UpdateUser(id int, user *models.Users) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	config.DB.First(&user, id)
	return user, nil
}

// function login database untuk mendapatkan token
func LoginUser(plan_pass string, user *models.Users) (interface{}, error) {
	var get_login models.GetLoginUser
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// cek plan password dengan hash password
	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plan_pass))
	if match != nil {
		return nil, match
	}
	user.Token, err = middlewares.CreateToken(int(user.ID)) // generate token
	if err != nil {
		return nil, err
	}
	get_login.ID = user.ID
	get_login.Name = user.Name
	get_login.Token = user.Token
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return get_login, nil
}
