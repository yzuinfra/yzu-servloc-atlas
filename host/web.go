package host

import (
	"yzuinfra/atlas/host/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupFiber(host string, port string) *fiber.App {
	app := fiber.New()
	routes.AddRegisteredHandler(app)
	app.Listen(host + ":" + port)
	return app
}

func RunWebHost(host string, port string) {
	go SetupFiber(host, port)
}
