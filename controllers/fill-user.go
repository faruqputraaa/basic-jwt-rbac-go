package controllers

import (
	"main/config"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func FillUser(c echo.Context) error {

	// Membuat user admin baru
	adminUser := models.User{
		Name:     "Admin1",
		Email:    "admin@mail.com",
		Password: "adminpassword1",
		Role:     "Admin",
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal melakukan hash password: " + err.Error(),
		})
	}
	adminUser.Password = string(hashedPassword)

	// Simpan ke database
	if err := config.DB.Create(&adminUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal membuat user: " + err.Error(),
		})
	}

	// Mengembalikan respon sukses
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Data Berhasil di buat",
		"user":    adminUser,
	})
}
