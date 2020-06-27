package main

import (
	"github.com/labstack/echo"
	"github.com/pekrj/labgolang/adapters/handlers"
)

func main() {
	h := handlers.Handler{}
	e := echo.New()

	e.GET("/courses", h.GetCourse)
	e.Logger.Fatal(e.Start(":1323"))
}
