package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article struct yapısı, JSON formatında serileştirme için alan etiketleri ile tanımlanmıştır.
type Article struct {
	Title  string `json:"Title"`
	Desc   string `json:"desc"`
	Contet string `json:"content"`
}

// Articles, Article tipinde bir dizi yapısını temsil eder.
type Articles []Article

// Tüm makaleleri döndüren handler fonksiyonu
func allArticals(w http.ResponseWriter, r *http.Request) {
	// Test amaçlı makaleler oluşturuldu
	articles := Articles{
		Article{Title: "Test Başlık", Desc: "Test Açıklama", Contet: "Merhaba Dünya"},
		Article{Title: "Test2 Başlık", Desc: "Test2 Açıklama", Contet: "Merhaba Dünya2"},
		Article{Title: "Test3 Başlık", Desc: "Test3 Açıklama", Contet: "Merhaba Dünya3"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint") // Konsola endpoint bilgisi yazdırılır
	json.NewEncoder(w).Encode(articles)                // Makaleleri JSON formatında HTTP yanıtına kodlar

	w.Write([]byte("Sadık Sünbül")) //boylede verı godnerılır ama byte seklınde gittigi içi
}

// Ana sayfa endpointi için handler fonksiyonu
func homePage(w http.ResponseWriter, r *http.Request) {
	// Tarayıcıya basit bir metin yanıtı gönderilir
	fmt.Fprintf(w, "Ana sayfa endpointi hit edildi")
}

// HTTP isteklerini yönetmek için handleRequest fonksiyonu
func handleRequest() {
	// "/" ve "/articles" endpointlerini ilgili handler fonksiyonlarına bağlar
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticals)

	// HTTP sunucusunu 8081 portunda başlatır ve log.Fatal ile hata durumunda programı durdurur
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Ana fonksiyon, HTTP istekleri yönetimini başlatır
func main() {
	handleRequest() // handleRequest fonksiyonunu çağırarak sunucuyu başlatır
}
