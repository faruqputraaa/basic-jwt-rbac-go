package main

import (
	"main/config"
	"main/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Hubungkan ke database
	config.ConnectDB()

	// Inisialisasi Echo
	e := echo.New()

	// Daftarkan routes
	routes.Init(e)

	// Jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}
