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

	route.Start(":" + os.Getenv("APP_PORT"))
}
