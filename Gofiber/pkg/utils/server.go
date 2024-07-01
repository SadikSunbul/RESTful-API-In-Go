package utils

import (
	"fmt"
	"github.com/SadikSunbul/RESTful-API-In-Go/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateServer(port int) {
	app := fiber.New()                              // Fiber web çerçevesi kullanarak yeni bir HTTP sunucusu oluşturur. Fiber, Go dilinde yazılmış hızlı ve minimalist bir web çerçevesidir.
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port))) //belirtilen port numarasında sunucuyu başlatır ve gelen istekleri dinlemeye başlar.
	//	fmt.Sprintf(":%d", port) ifadesi, port numarasını bir stringe dönüştürür ve ":port" formatında bir string oluşturur. Örneğin, port 8080 ise, ":8080" stringi oluşturulur.
	//log.Fatal fonksiyonu, app.Listen fonksiyonunun döndürdüğü hata varsa, bu hatayı loglar ve programı sonlandırır. Bu, sunucu başlatılamazsa programın hata mesajıyla birlikte kapanmasını sağlar.
	routes.PostRoutes(app)
}
