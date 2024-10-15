package controllers

import (
	"main/config"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	// Validasi apakah data JSON yang dikirim sesuai dengan struktur User
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return c.JSON(http.StatusBadRequest, "Error: Data JSON tidak sesuai dengan struktur User")
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: Gagal menghash password: "+err.Error())
	}
	user.Password = string(hashedPassword)

	config.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	var users []models.User
	id := c.Param("id")
	config.DB.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User dengan ID " + id + " berhasil diupdate",
		"data":    users,
	})
}

func GetUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error: User tidak ditemukan")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User dengan ID " + id + " berhasil ditemukan",
		"data":    user,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error: User tidak ditemukan")
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	// Validasi apakah data JSON yang dikirim sesuai dengan struktur User
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return c.JSON(http.StatusBadRequest, "Error: Data JSON tidak sesuai dengan struktur User")
	}

	config.DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User dengan ID " + id + " berhasil diupdate",
		"data":    user,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user models.Todo
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Todo tidak ditemukan")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal menghapus Todo: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User dengan ID " + id + " berhasil dihapus",
		"data":    user,
	})
}
