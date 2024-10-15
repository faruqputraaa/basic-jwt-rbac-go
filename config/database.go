package config

import (
	"fmt"
	"log"
	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Sesuaikan dengan konfigurasi MySQL
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3333)/todo_db?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}

	fmt.Println("Database terhubung dengan sukses")


	// Jalankan AutoMigrate untuk membuat tabel
	DB.AutoMigrate(&models.User{}, &models.Todo{})
	fmt.Println("Database berhasil dimigrasi")
}
