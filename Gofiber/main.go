package main

import (
	"github.com/SadikSunbul/RESTful-API-In-Go/pkg/utils"
	"github.com/SadikSunbul/RESTful-API-In-Go/platform/migrations"
	"github.com/gofiber/fiber/v2"
)

/*
GoFiber :

	Gofiber framework kullanmamızın sebebini şu şekilde listeleyebiliriz:

	Yüksek performans ve düşük bellek kullanımı
	Hızlı server-side programlama
	Kütüphane tarafından sağlanan middlewarelar
	Kolay rotalama
*/
func main() {
	utils.CreateServer(3000)
	if !fiber.IsChild() {
		// Application initialization
		migrations.Migrate()
	}
}
