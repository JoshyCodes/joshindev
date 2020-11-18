package public

import (
	"github.com/gofiber/fiber/v2"
)

func LoadPublicHandlers(app *fiber.App) *fiber.App {

    app.Get("/", func(c *fiber.Ctx) error {
        // Render index template
        return c.Render("index", fiber.Map{})
	})
	
	return app
}


