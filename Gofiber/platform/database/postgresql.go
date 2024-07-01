package database

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
	"os"
)

// Config yapılandırma dosyası
type Config struct {
	Database struct { // veritabanı bilgileri
		Connection string `yaml:"connection"` // veritabanı bağlantısı
	} `yaml:"database"` // veritabanı bilgileri
}

var Conn *gorm.DB

func initializePostgres() *gorm.DB {
	var err error

	configFile, err := os.ReadFile("config.yaml") // config dosyasını okur
	if err != nil {
		log.Fatalf("config.yaml dosyası okunamadı: %v", err) // hata döndürür
	}

	var config Config                         // yapılandırma dosyası
	err = yaml.Unmarshal(configFile, &config) // yapılandırma dosyasını okur
	if err != nil {                           // hata varsa
		log.Fatalf("Yapılandırma dosyası ayrıştırılamadı: %v", err)
	}

	conn, _ := url.Parse(config.Database.Connection)       // veritabanı bağlantısı
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem" // ssl bilgileri

	Conn, err = gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
	fmt.Println("-----------------:", conn.String())
	if err != nil {
		panic(err)
	}

	return Conn
}
