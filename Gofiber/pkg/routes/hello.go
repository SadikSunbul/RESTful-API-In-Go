package routes

import (
	"github.com/SadikSunbul/RESTful-API-In-Go/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func HelloRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
