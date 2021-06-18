package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pudjamansyurin/golang/echo-gorm/models"
)

func GetUsers(c echo.Context) error {
	user := new(models.User)
	return c.JSON(http.StatusOK, user)
}
