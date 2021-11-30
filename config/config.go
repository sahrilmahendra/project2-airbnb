package config

import (
	"log"
	"os"

	"project2/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// inisialisasi database
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := os.Getenv("CONNECTION_DB")

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// auto migrate -> untuk membuat tabel otomatis jika tabel tidak terdapat pada database
func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Reservation{})
	DB.AutoMigrate(&models.CreditCard{})
	DB.AutoMigrate(&models.Homestay{})
	DB.AutoMigrate(&models.Facility{})
	DB.AutoMigrate(&models.Homestay_Facility{})
}

// ===============================================================//

// inisialisasi database untuk untuk unit testing
func InitDBTest() {
	config_testing := os.Getenv("CONNECTION_DB_TESTING")

	var e error
	DB, e = gorm.Open(mysql.Open(config_testing), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrationTest()
}

// auto migrate -> untuk membuat tabel otomatis jika tabel tidak terdapat pada database
// drop table -> untuk menghapus tabel terlebih dahulu agar isi datanya dimulai dari tabel kosong
func InitMigrationTest() {
	DB.Migrator().DropTable(&models.Users{})
	DB.Migrator().DropTable(&models.Reservation{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Reservation{})
}
