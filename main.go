package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", rootHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
	log.Println("This is silly. Convert it into something great!")
}

func rootHandler(ctx *gin.Context) {
	dbEndpoint := os.Getenv("DB_ENDPOINT")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	output := "DB_ENDPOINT: " + dbEndpoint + "\n" +
						"DB_PORT: " + dbPort + "\n" +
						"DB_NAME: " + dbName + "\n" + 
						"This is a silly demo"

	ctx.String(http.StatusOK, output)
}
