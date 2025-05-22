package main

import (
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    app.Use(cors.New())

    app.Get("/api/tasks", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello from Go API"})
    })

    port := "8080"
    fmt.Println("Listening on port", port)
    log.Fatal(app.Listen(":" + port))
}
