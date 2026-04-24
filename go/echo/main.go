package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
func pingHandler(c *echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

/* ---------- /json ---------- */

type StatusResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp int64  `json:"timestamp"`
}

func jsonHandler(c *echo.Context) error {
	resp := StatusResponse{
		Status:    "ok",
		Service:   "echo",
		Timestamp: time.Now().Unix(),
	}
	return c.JSON(http.StatusOK, resp)
}

/* ---------- /echo ---------- */

type EchoMessage struct {
	Message string `json:"message"`
}

func echoHandler(c *echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "expected application/json"},
		)
	}

	var payload EchoMessage
	if err := c.Bind(&payload); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "invalid json"},
		)
	}

	return c.JSON(http.StatusOK, payload)
}

/* ---------- main ---------- */

func main() {
	e := echo.New()
	e.GET("/ping", pingHandler)
	e.GET("/json", jsonHandler)
	e.POST("/echo", echoHandler)

	if err := e.Start("127.0.0.1:8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
