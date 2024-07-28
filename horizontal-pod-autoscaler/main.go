package main

import (
	"fmt"
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routes
	app.Get("/cpuLoad", cpuHandler)
	app.Get("/memoryLoad", memoryHandler)
	app.Get("/healthz", status)
	app.Get("/readyz", status)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

// Handlers
func cpuHandler(c *fiber.Ctx) error {
	x := 100.0
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	return c.SendString("SQRT of 0.0001: " + fmt.Sprint(x) + "\n")
}

func memoryHandler(c *fiber.Ctx) error {
	var dummyArray [1000]int
	fmt.Println(dummyArray)
	return c.SendString("Allocated 8KB\n")
}

func status(c *fiber.Ctx) error {
	return c.SendString("OK!")
}
