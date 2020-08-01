package handlers

import (
	"net/http"

	"github.com/labstack/gommon/log"

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

	log.Info("Passei No handler e chamei o Get Courses")
	return c.JSON(http.StatusOK, courses)
}
