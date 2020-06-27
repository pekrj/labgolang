package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct{}

func (h *handler) getCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
