package post

import (
	"github.com/SadikSunbul/RESTful-API-In-Go/app/models"
	"github.com/SadikSunbul/RESTful-API-In-Go/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	database.Conn.Find(&posts)

	return ctx.JSON(posts)
}

/*
GORM ile bir modelin girdilerini listelemek istediğimizde yukarıdaki yola başvurabiliriz. Find fonksiyonu bizim için
tüm postları getirecektir ve bunu da kullanıcılara JSON olarak döndürebiliriz.
*/
