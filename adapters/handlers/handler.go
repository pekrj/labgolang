package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pekrj/labgolang/domain/services"
)

//Handler to handle requests
type Handler struct {
	CourseService services.CourseService
}

//GetCourses get courses
func (h *Handler) GetCourses(c echo.Context) error {
	courses, err := h.CourseService.GetAllCourses()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, courses)
}
