package public

import (
	"joshindev/internal/handlers/public"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func CreatePublicRouter() *fiber.App {
	
	engine := html.New("website/templates", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
	})
	
	return public.LoadPublicHandlers(&app)
}