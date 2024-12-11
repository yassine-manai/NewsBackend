package main

import (
	"log"
	"news-app/internal/api"
	"news-app/internal/db"
	"news-app/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.DB.Close()

	router := gin.Default()

	api.RegisterRoutes(router)
	web.ServeFrontend(router)

	port := ":8070"
	log.Printf("Server running on %s", port)
	log.Fatal(router.Run(port))
}
