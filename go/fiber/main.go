package main 

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
func pingHandler(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/plain")
	return c.SendString("ok")
}

/* ---------- /json ---------- */

type StatusResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp int64  `json:"timestamp"`
}

func jsonHandler(c *fiber.Ctx) error {
	resp := StatusResponse{
		Status:    "ok",
		Service:   "fiber",
		Timestamp: time.Now().Unix(),
	}

	return c.JSON(resp)
}

/* ---------- /echo ---------- */

type EchoMessage struct {
	Message string `json:"message"`
}

func echoHandler(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "expected application/json"})
	}

	var payload EchoMessage
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid json"})
	}

	return c.JSON(payload)
}

/* ---------- main ---------- */

func main() {
	// Fiber is already "release" by default
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/ping", pingHandler)
	app.Get("/json", jsonHandler)
	app.Post("/echo", echoHandler)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}
