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

func AddRegisteredHandler(app *fiber.App) {
	app.Get("/registered", RegisteredHandler)
}
