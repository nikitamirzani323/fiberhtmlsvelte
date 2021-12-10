package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func Init() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/review/*", func(c *fiber.Ctx) error {
		// Render index template
		urlsegment := c.Params("*")
		log.Println(urlsegment)
		return c.Render("index", fiber.Map{
			"Title":      "Home",
			"Page":       "Home",
			"Urlsegment": urlsegment,
		})
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		// Render index template
		urlsegment := c.Params("*")
		log.Println(urlsegment)
		return c.Render("about", fiber.Map{
			"Title": "About",
			"Page":  "About",
		})
	})
	return app
}
