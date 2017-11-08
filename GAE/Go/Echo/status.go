package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	response struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Data    []string `json:"data"`
	}
)

func init() {
	// hook into the echo instance to create an endpoint group
	// and add specific middleware to it plus handlers
	g := e.Group("/status")
	g.Use(middleware.CORS())

	g.GET("/:statusCode", httpResponse)
}

// httpResponse エラー時のレスポンスを返す
func httpResponse(c echo.Context) (err error) {
	statusCode, err := strconv.Atoi(c.Param("statusCode"))
	if err != nil {
		statusCode = http.StatusNotFound
	}

	// 存在しないstatus_codeが来た場合は404
	message := http.StatusText(statusCode)
	if message == "" {
		statusCode = http.StatusNotFound
		message = http.StatusText(http.StatusNotFound)
	}

	response := response{Status: statusCode, Message: message, Data: make([]string, 0)} // make([]string, 0)で空の配列を生成
	return c.JSON(statusCode, response)
}
