package routes

import (
	"github.com/SadikSunbul/RESTful-API-In-Go/app/controllers/post"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	// List All Posts
	app.Get("/posts", post.Index)

	// Create Post
	app.Post("/post", post.Create)
}

/*
Bu PostRoutes fonksiyonumuzu da CreateServer altında fiber app oluşturduktan hemen sonra çağırabiliriz.
routes.PostRoutes(app) şeklinde tanımlamayı yaptığımızda rotalarımız kullanıma hazır olacaktır.
*/
