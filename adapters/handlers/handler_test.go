package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetCourses(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/courses", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(req, response)
	h := &handler{}

	if assert.NoError(t, h.getCourses(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}

}
