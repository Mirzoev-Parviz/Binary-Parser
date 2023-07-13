package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadEnv() (*DatabaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbConfig := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	return dbConfig, nil
}

var DB = ConnectDB()

/*
func ConnectDB() *sql.DB {
	dbConfig, err := LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	return db
}
*/

func ConnectDB() *sql.DB {
	// Замените значения на свои параметры подключения
	info, err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.Username, info.Password, info.Database)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("ошибка подключения к базе данных: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("ошибка проверки подключения к базе данных: %s", err.Error())
	}

	return db
}

func DisconnectDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

/*
package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadEnv() (*DatabaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbConfig := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	return dbConfig, nil
}

var DB = ConnectDB()

func ConnectDB() *sql.DB {
	dbConfig, err := LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password,
		dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	return db
}

func DisconnectDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
*/
