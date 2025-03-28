package routes

import (
	"encoding/json"
	"yzuinfra/atlas/host/regmap"

	"github.com/gofiber/fiber/v2"
)

func RegisteredHandler(c *fiber.Ctx) error {
	data, err := json.Marshal(regmap.GetAgents())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(string(data))
}

func RegisteredWithServiceHandler(c *fiber.Ctx) error {
	service := c.Params("service")
	data, err := json.Marshal(regmap.GetAgentsWithService(service))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(string(data))
}

func RegisteredWithHostnameHandler(c *fiber.Ctx) error {
	hostname := c.Params("hostname")
	data, err := json.Marshal(regmap.GetAgentsFromHostname(hostname))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(string(data))
}

func AddRegisteredHandler(app *fiber.App) {
	app.Get("/registered", RegisteredHandler)
	app.Get("/registered/service/:service", RegisteredWithServiceHandler)
	app.Get("/registered/hostname/:hostname", RegisteredWithHostnameHandler)
}
