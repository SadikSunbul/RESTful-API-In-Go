package migrations

import (
	"github.com/SadikSunbul/RESTful-API-In-Go/app/models"
	"github.com/SadikSunbul/RESTful-API-In-Go/platform/database"
)

func Migrate() {
	database.Conn.AutoMigrate(&models.Post{})
}

/*
Migration işlemi oluşturduğumuz modellerin otomatik şekilde veritabanında tablolarının eklenmesi işlemidir.
Bunun için platform/migrations altında migrate.go isminde bir dosya oluşturalım ve içeriğini aşağıdaki şekilde girelim.

Bu kısımdaki AutoMigrate fonksiyonu GORM’un içerisinde gelen bir fonksiyondur. Diğer oluşturduğunuz modeller (varsa) bu
kısımda tanımlayarak uygulama başlangıcında tablolarının oluşturulmasını sağlayabilirsiniz. Bu fonksiyonu CreateServer
	fonksiyonu altında database.Init fonksiyonundan hemen sonra çağırabiliriz.
*/
