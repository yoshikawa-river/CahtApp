package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	HttpInfo *HttpInfo
	DBInfo   *DBInfo
}

type HttpInfo struct {
	PROTOCOL string
}

type DBInfo struct {
	CONN       string
	USER       string
	PASSWORD   string
	DBPROTOCOL string
	DBNAME     string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	protocol := os.Getenv("PORT")

	httpInfo := &HttpInfo{
		PROTOCOL: protocol,
	}

	conn := os.Getenv("DB_CONNECTION")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	DBProtocol := os.Getenv("DB_PROTOCOL")
	dbname := os.Getenv("DB_NAME")

	dbInfo := &DBInfo{
		CONN:       conn,
		USER:       user,
		PASSWORD:   password,
		DBPROTOCOL: DBProtocol,
		DBNAME:     dbname,
	}

	return &AppConfig{
		HttpInfo: httpInfo,
		DBInfo:   dbInfo,
	}
}
