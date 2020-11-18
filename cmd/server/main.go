package main

import (
	"log"

	"joshindev/internal/handlers/public"

	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2"
)

func main() {

    engine := html.New("website/templates", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
	})

	app = public.LoadPublicHandlers(app)

	log.Fatal(app.Listen(":8081"))
}