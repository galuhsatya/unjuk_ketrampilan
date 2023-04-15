package controller

import (
	"net/http"
	"prakerja2/config"
	"prakerja2/model"

	"github.com/labstack/echo/v4"
)

func AddProfiles(c echo.Context) error {

	var profile model.Profile
	c.Bind(&profile)

	result := config.DB.Create(&profile)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: profile,
	})
}

func GetDetailProfiles(c echo.Context) error {
	id := c.Param("id")
	// logic
	// db.Where("name = ?", "jinzhu").First(&user)
	var profile model.Profile
	result := config.DB.Where("id = ?", id).First(&profile)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: profile,
	})
}

func GetProfiles(c echo.Context) error {
	var profiles []model.Profile

	result := config.DB.Find(&profiles)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: profiles,
	})
}

func DeleteProfile(c echo.Context) error {
	id := c.Param("id")
	// logic
	// db.Where("name = ?", "jinzhu").Delete(&email)
	var profile model.Profile
	result := config.DB.Where("id = ?", id).Delete(&profile)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: profile,
	})
}
