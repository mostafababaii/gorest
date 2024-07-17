package app

import (
	"context"
	"fmt"
	"github.com/mostafababaii/gorest/internal/database/mysql"
	"github.com/mostafababaii/gorest/internal/database/redis"
	"github.com/mostafababaii/gorest/internal/models"
	"github.com/mostafababaii/gorest/internal/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mostafababaii/gorest/config"
)

func init() {
	config.Setup()
	mysql.Setup()
	redis.Setup()

	models.Migrate(mysql.NewConnection())
}

func Start() {
	gin.SetMode(config.ServerConfig.RunMode)
	routersInit := routers.InitRouter()
	serverAddress := fmt.Sprintf(":%d", config.ServerConfig.HttpPort)

	server := &http.Server{
		Addr:           serverAddress,
		Handler:        routersInit,
		ReadTimeout:    config.ServerConfig.ReadTimeout,
		WriteTimeout:   config.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // One MegaByte
	}

	go func() {
		log.Printf("Starting server on port: %d", config.ServerConfig.HttpPort)

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Receive terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(tc)
}
