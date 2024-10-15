package routes

import (
	"main/controllers"
	"main/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	e.POST("/login", controllers.Login)
	e.POST("/fill-user", controllers.FillUser)

	// Routes untuk Todo-list, hanya bisa diakses oleh pengguna dengan role Editor
	todoRoutes := e.Group("/todos")
	todoRoutes.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("Editor"))
	todoRoutes.POST("", controllers.CreateTodo)
	todoRoutes.GET("", controllers.GetTodos)
    todoRoutes.GET("/:id", controllers.GetTodo)
	todoRoutes.PUT("/:id", controllers.UpdateTodo)
	todoRoutes.DELETE("/:id", controllers.DeleteTodo)

	// Routes untuk User, hanya bisa diakses oleh pengguna dengan role Admin
	userRoutes := e.Group("/users")
	userRoutes.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("Admin"))
	userRoutes.POST("", controllers.CreateUser)
	userRoutes.GET("", controllers.GetUsers)
	userRoutes.GET("/:id", controllers.GetUser)
	userRoutes.PUT("/:id", controllers.UpdateUser)
	userRoutes.DELETE("/:id", controllers.DeleteUser)

}
