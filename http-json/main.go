package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type superhero struct {
	Name     string   `json:"name"`
	Age      uint     `json:"age"`
	Height   float32  `json:"height"`
	Aliases  []string `json:"aliases"`
	Location location `json:"location"`
}

type location struct {
	Street string `json:"street"`
	Number uint   `json:"number"`
	Hidden bool   `json:"hidden"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		batman := &superhero{
			Name:   "Bruce Wayne",
			Age:    38,
			Height: 1.879,
			Aliases: []string{
				"Batman", "The Dark Knight",
			},
			Location: location{
				Street: "Mountain Drive",
				Number: 1007,
				Hidden: false,
			},
		}

		return c.Status(fiber.StatusOK).JSON(batman)
	})

	if err := app.Listen(":3001"); err != nil {
		log.Errorf("process ended with error: ")
	}
}
