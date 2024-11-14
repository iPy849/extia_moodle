package app

import (
	"extia/app/handlers"
	"extia/logger"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/mustache/v2"
)

func RunApp() {
	log.SetOutput(logger.Logger.Writer())

	engine := mustache.New("./views", ".mustache")

	app := fiber.New(fiber.Config{
		AppName:        fmt.Sprintf("%s v0.0.0", os.Getenv("APP_NAME")),
		ReadTimeout:    10 * time.Second,       // Timeout
		WriteTimeout:   500 * time.Millisecond, // Timeout
		ReadBufferSize: 1 << 20,                // Maximum message length
		Views:          engine,                 // Templating engine
		ViewsLayout:    "layouts/main",         // Templating engine default layout
	})

	BindHandlersToApp(
		app,
		new(handlers.RenderHandler),
		new(handlers.APIHandler),
	)

	log.Info("Listening on ", os.Getenv("APP_HOST"))
	log.Fatal(app.Listen(os.Getenv("APP_HOST")))
}
