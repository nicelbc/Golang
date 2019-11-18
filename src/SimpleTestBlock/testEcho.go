package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func tmain() {
	// Echo instance 框架开启
	e := echo.New()

	// Middleware 日志，回复
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/go", hello2)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler，处理跳转
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func hello2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! The second hello")
}
