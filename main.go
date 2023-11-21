package main

import (
	"log"
	"os"

	"github.com/fazaalexander/montirku-be/cmd/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	log.Println("Starting application...")
	route := app.StartApp()

	if err := route.Start(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}
}
