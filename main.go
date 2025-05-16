package main

import (
	"BlogHub/pkg/db"
	"BlogHub/pkg/routers"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {

	db.Connect()
	router := routers.Routes()

	router.Run()
}
