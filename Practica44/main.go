package main

import (
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.POST("/persons", save)
	e.GET("/persons", FindAllPersons)
	e.GET("/persons/:id", FindPersonByID)
	e.PUT("/persons/:id", UpdateAUser)
	e.DELETE("/persons/:id", DeletePerson)
	handleRequest()
}
