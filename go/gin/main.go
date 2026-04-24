package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
func pingHandler(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"text/plain",
		[]byte("ok"),
	)
}

/* ---------- /json ---------- */

type StatusResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp int64  `json:"timestamp"`
}

func jsonHandler(c *gin.Context) {
	resp := StatusResponse{
		Status:    "ok",
		Service:   "gin",
		Timestamp: time.Now().Unix(),
	}

	c.JSON(http.StatusOK, resp)
}

/* ---------- /echo ---------- */

type EchoMessage struct {
	Message string `json:"message"`
}

func echoHandler(c *gin.Context) {
	// Enforce Content-Type = application/json
	if c.GetHeader("Content-Type") != "application/json" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "expected application/json"},
		)
		return
	}

	var payload EchoMessage
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid json"},
		)
		return
	}

	c.JSON(http.StatusOK, payload)
}

/* ---------- main ---------- */

func main() {
	// IMPORTANT for benchmarking:
	// gin.Default() includes logger + recovery (adds overhead)
	// Use gin.New() for fair comparison
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/ping", pingHandler)
	r.GET("/json", jsonHandler)
	r.POST("/echo", echoHandler)

	// Match net/http bind address
	r.Run("127.0.0.1:8080")
}
