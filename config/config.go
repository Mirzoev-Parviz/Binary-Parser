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

func ConnectDB() *sql.DB {
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

func GetDBConn() *sql.DB {
	return DB
}

func DisconnectDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
