package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Handler to handle requests
type Handler struct{}

//GetCourses get courses
func (h *Handler) GetCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, "Funcionando")
}
