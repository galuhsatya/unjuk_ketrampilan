package route

import (
	"prakerja2/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/profiles", controller.AddProfiles)
	e.GET("/profiles", controller.GetProfiles)
	e.GET("/profiles/:id", controller.GetDetailProfiles)
	e.DELETE("/profiles/:id", controller.DeleteProfile)

	return e
}
