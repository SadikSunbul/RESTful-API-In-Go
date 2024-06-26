package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

	/*
		w.Write([]byte("Sadık Sünbül \n")) //boylede verı godnerılır ama json seklınde gitmediği için kullanmayız

		deneme := Deneme{ //istersen altakı gııbde godner bırlrız verıleris
			Name:  "John",
			Name2: "Doe",
		}
		jsonData, err := json.Marshal(deneme)
		if err != nil {
			fmt.Println("JSON encode hatası:", err)
			return
		}

		w.Write(jsonData)

		// Buda olur
		json.NewEncoder(w).Encode(struct {
			Name  string `json:"ali"` //bu degıskenın karsılıgı alı olyur jsonda
			Name2 string `json:"salih"`
		}{
			Name:  "Ahmet", //degiri burada verilir
			Name2: "Mehmet",
		}) */

}

//type Deneme struct {
//	Name  string `json:"ali"`
//	Name2 string `json:"salih"`
//}

// Ana sayfa endpointi için handler fonksiyonu
func homePage(w http.ResponseWriter, r *http.Request) {
	// Tarayıcıya basit bir metin yanıtı gönderilir
	fmt.Fprintf(w, "Ana sayfa endpointi hit edildi")
}

func testPostArticals(w http.ResponseWriter, r *http.Request) {
	// Tarayıcıya basit bir metin yanıtı gönderilir
	fmt.Fprintf(w, "testPostArticals endpointi hit edildi")
}

// HTTP isteklerini yönetmek için handleRequest fonksiyonu
func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	// "/" ve "/articles" endpointlerini ilgili handler fonksiyonlarına bağlar
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", allArticals)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", allArticals).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticals).Methods("POST")

	// HTTP sunucusunu 8081 portunda başlatır ve log.Fatal ile hata durumunda programı durdurur
	log.Fatal(http.ListenAndServe(":8081", myRouter)) //2.parametre nildi onceden
}

// Ana fonksiyon, HTTP istekleri yönetimini başlatır
func main() {
	handleRequest() // handleRequest fonksiyonunu çağırarak sunucuyu başlatır
}
