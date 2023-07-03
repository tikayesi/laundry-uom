package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	DataSourceName string
}

func (c *Config) initDb() {
	// tampung nilai ENV dari terminal
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "tikayesikristiani"
	// dbPass := ``
	// dbName := "laundry_db"

	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	// buka koneksi
	c.DataSourceName = dsn
}

func NewConfig() Config {
	config := Config{}
	config.initDb() // untuk menjalankan method yang didalamnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang didalamnya terdapat attribute bertipe data Koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
}
