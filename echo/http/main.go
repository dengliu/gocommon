package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// https://echo.labstack.com/guide/

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)

	// Initialize handler
	h := Handler{client: http.Client{Timeout: time.Second * 10}}

	// Routes
	e.GET("/", h.get)
	e.POST("/users", h.post)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type JSONResponse struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler

type Handler struct {
	client http.Client
}

func (h *Handler) get(c echo.Context) error {
	resp := JSONResponse{
		Data: user{
			Name: "Joe Biden",
			Age:  75,
		},
		Error: "fake error",
	}
	return c.JSON(http.StatusOK, resp)
}

// curl -v -X POST http://localhost:8080/users  -H 'Content-Type: application/json'   -d '{"name":"Adam Smith","age": 90}'

func (h *Handler) post(c echo.Context) error {
	u := new(user)
	// 	u := new(map[string]any)
	if err := c.Bind(u); err != nil {
		return err
	}

	c.Logger().Debug("==== this is a debug message =====")
	return c.JSON(http.StatusOK, JSONResponse{
		Data:  u,
		Error: "biz logic error",
	})
}
