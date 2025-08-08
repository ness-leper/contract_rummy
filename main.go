package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Static("./templates/", "templates")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/index.tmpl")
	})

	app.Get("/round/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		return c.JSON(fiber.Map{"message": "We were here " + idStr})
	})

	app.Get("/newPlayer", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/newPlayer.tmpl")
	})

	app.Get("/newGame", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/setup.tmpl")
	})

	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"id":   id,
			"name": "John Doe",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
