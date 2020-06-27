package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) GetCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
