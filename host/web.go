package host

import (
	"yzuinfra/atlas/host/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupFiber() *fiber.App {
	app := fiber.New()
	routes.AddRegisteredHandler(app)
	app.Listen(":8000")
	return app
}

func RunWebHost() {
	SetupFiber()
}
