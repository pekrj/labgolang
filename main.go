package main

import (
	"github.com/labstack/echo"
	"github.com/pekrj/labgolang/adapters/handlers"
	"github.com/pekrj/labgolang/adapters/inmemorycourseRepository"
	"github.com/pekrj/labgolang/domain/services"
)

func main() {
	h := &handlers.Handler{
		CourseService: services.CourseService{
			CourseRepository: &inmemorycourseRepository.InMemoryCourseRepository{},
		},
	}
	e := echo.New()

	e.GET("/courses", h.GetCourses)
	e.Logger.Fatal(e.Start(":1323"))
}
