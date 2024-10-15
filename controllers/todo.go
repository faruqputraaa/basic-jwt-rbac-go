package controllers

import (
	"main/config"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	// Ambil userID dari token JWT
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusBadRequest, "userID tidak valid")
	}

	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Set userID dari user yang sedang login
	todo.UserID = userID

	// Validasi title dan context
	if todo.Title == "" || todo.Content == "" {
		return c.JSON(http.StatusBadRequest, "Judul dan Konten tidak boleh kosong")
	}

	// Simpan Todo ke database
	if err := config.DB.Create(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	config.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error: Data tidak ditemukan")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Todo dengan ID " + id + " berhasil ditemukan",
		"data":    todo,
	})
}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo

	// Cek apakah Todo dengan ID yang diberikan ada di database
	if err := config.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Todo tidak ditemukan",
		})
	}

	// Bind JSON dari request ke struct Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error saat parsing data JSON: " + err.Error(),
		})
	}

	// Validasi apakah semua field yang diperlukan tersedia dan valid
	if todo.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Data 'Title' tidak boleh kosong",
		})
	}
	if todo.Content == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Data 'Content' tidak boleh kosong",
		})
	}

	// Jika semua validasi berhasil, simpan perubahan ke database
	if err := config.DB.Save(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal menyimpan perubahan ke database: " + err.Error(),
		})
	}

	// Berikan respon sukses
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Todo berhasil diupdate",
		"data":    todo,
	})
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Todo tidak ditemukan")
	}

	if err := config.DB.Delete(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal menghapus Todo: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Todo berhasil dihapus",
		"data":    todo,
	})
}
