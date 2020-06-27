package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pekrj/labgolang/domain/services"

	"github.com/pekrj/labgolang/adapters/inmemorycourseRepository"
	"github.com/pekrj/labgolang/domain"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetCourses(t *testing.T) {
	var inMemoryRepository = &inmemorycourseRepository.InMemoryCourseRepository{}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/courses", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(req, response)
	h := &Handler{
		CourseService: services.CourseService{
			CourseRepository: inMemoryRepository,
		},
	}

	if assert.NoError(t, h.GetCourses(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
		var courses []domain.Course
		json.Unmarshal([]byte(response.Body.String()), &courses)
		expectedCourses, _ := inMemoryRepository.GetAll()
		assert.Equal(t, expectedCourses, courses)
	}

}
