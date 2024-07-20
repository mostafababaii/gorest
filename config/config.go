package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	ServerConfig     server
	DatabaseConfig   database
	RedisConfig      redis
	JwtSecret        []byte
	JwtTokenLifeSpan time.Duration
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Driver      string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	MaxIdle     int
	MaxOpen     int
}

type redis struct {
	Host        string
	Port        int
	User        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
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

	dbMaxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	if err != nil {
		dbMaxIdle = 10
	}

	dbMaxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	if err != nil {
		dbMaxOpen = 100
	}

	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		redisPort = 6379
	}

	redisMaxIdle, err := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	if err != nil {
		redisMaxIdle = 10
	}

	redisMaxActive, err := strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE"))
	if err != nil {
		redisMaxActive = 100
	}

	IdleTimeoutInSecond, err := strconv.Atoi(os.Getenv("REDIS_IDLE_TIMEOUT_IN_SECOND"))
	if err != nil {
		IdleTimeoutInSecond = 300
	}

	ServerConfig = server{
		RunMode:      os.Getenv("RUN_MODE"),
		HttpPort:     serverPort,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(WriteTimeout) * time.Second,
	}

	DatabaseConfig = database{
		Driver:      os.Getenv("DB_DRIVER"),
		User:        os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		Host:        os.Getenv("DB_HOST"),
		Name:        os.Getenv("DB_NAME"),
		TablePrefix: os.Getenv("DB_TABLE_PREFIX"),
		MaxIdle:     dbMaxIdle,
		MaxOpen:     dbMaxOpen,
	}

	RedisConfig = redis{
		Host:        os.Getenv("REDIS_HOST"),
		Port:        redisPort,
		User:        os.Getenv("REDIS_USER"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		MaxIdle:     redisMaxIdle,
		MaxActive:   redisMaxActive,
		IdleTimeout: time.Duration(IdleTimeoutInSecond) * time.Second,
	}
}
