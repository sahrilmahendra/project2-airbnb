package routes

import (
	"project2/constants"
	"project2/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// route users tanpa JWT
	e.POST("/users", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUserControllers)
	e.POST("/reservation", controllers.CreateReservationControlllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// route users dengan JWT
	j.GET("/users/:id", controllers.GetUserControllers)
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	// route homestay dengan JWT
	j.POST("/homestay", controllers.CreateHomestayControllers)
	j.GET("/homestay", controllers.GetAllHomestayControllers)
	j.GET("/homestay/:id", controllers.GetHomestayByIdControllers)
	j.PUT("/homestay/:id", controllers.UpdateHomestayControllers)
	j.DELETE("/homestay/:id", controllers.DeleteHomestayControllers)
	return e

}
