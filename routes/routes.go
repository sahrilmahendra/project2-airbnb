package routes

import (
	"project2/constants"
	"project2/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// users
	e.POST("/users", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUserControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// users
	j.GET("/users/:id", controllers.GetUserControllers)
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	// homestay
	j.POST("/homestay", controllers.CreateHomestayControllers)
	return e

}
