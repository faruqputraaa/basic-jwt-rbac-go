package controllers

import (
	"main/config"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var loginRequest LoginRequest

	// Bind JSON ke struct LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	// Ambil data user dari database
	var user models.User
	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	// Cek apakah password yang diinput cocok dengan password yang di-hash di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generate token dan response
	token, err := config.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
