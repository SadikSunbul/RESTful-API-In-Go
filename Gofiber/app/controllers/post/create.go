package post

import (
	"github.com/SadikSunbul/RESTful-API-In-Go/app/models"
	"github.com/SadikSunbul/RESTful-API-In-Go/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}

	if err := ctx.BodyParser(&post); err != nil { //Öncelikle Fiber’ın BodyParser fonksiyonunu kullanarak bodyden gelen JSON’u parse edelim.
		return ctx.Status(503).JSON(err)
	}
	if len(post.Content) < 1 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "İçerik gerekli",
		})
	}

	if err := database.Conn.Create(&post).Error; err != nil { //Ardından database.Conn.Create diyerek girdimizi veritabanına ekleyelim.
		return ctx.Status(503).JSON(err)
	}

	return ctx.JSON(post) // İşlemlerin sonunda da kullanıcıya JSON olarak oluşturulan gönderimizi gönderebiliriz.

}
