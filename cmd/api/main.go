package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-money-transfer/config"
	"github.com/go-money-transfer/internal/database"
	"github.com/go-money-transfer/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c, err := config.Init()
	if err != nil {
		log.Fatalf("error while connecting to database: %s", err)
	}

	err = database.Connect(c)
	if err != nil {
		log.Fatalf("error while connecting to database: %s", err)
	}
	r := fiber.New()
	router.SetupRoutes(r)

	go r.Listen(os.Getenv("PORT"))
	setGracefulShutdown()
}

func setGracefulShutdown() {
	n := make(chan os.Signal)
	signal.Notify(n, syscall.SIGINT, syscall.SIGTERM)

	<-n
	fmt.Println("Server is shutting down...")
	time.Sleep(time.Millisecond * 500)
}
