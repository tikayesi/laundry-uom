package config

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	DataSourceName string
}

func (c *Config) initDb() {
	// tampung nilai ENV dari terminal
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "tikayesikristiani"
	dbPass := ``
	dbName := "laundry_db"

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
