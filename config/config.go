package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	DatabaseConfig   *database
	ServerConfig     *server
	JwtSecret        []byte
	JwtTokenLifeSpan time.Duration
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

type database struct {
	Driver      string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Setup() {
	authSecret := os.Getenv("JWT_SECRET_KEY")
	if authSecret == "" {
		authSecret = "app_secret_key"
	}

	JwtSecret = []byte(authSecret)

	tokenLifeSpan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		tokenLifeSpan = 1
	}

	JwtTokenLifeSpan = time.Duration(tokenLifeSpan) * time.Hour

	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		readTimeout = 2
	}

	WriteTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		WriteTimeout = 2
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		serverPort = 8080
	}

	ServerConfig = &server{
		RunMode:      os.Getenv("RUN_MODE"),
		HttpPort:     serverPort,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(WriteTimeout) * time.Second,
	}

	DatabaseConfig = &database{
		Driver:      os.Getenv("DB_DRIVER"),
		User:        os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		Host:        os.Getenv("DB_HOST"),
		Name:        os.Getenv("DB_NAME"),
		TablePrefix: os.Getenv("DB_TABLE_PREFIX"),
	}
}
