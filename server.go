package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var uploadDest string = "data"

func handleRoot(c *gin.Context) {
	// If the client is 192.168.1.2, use the X-Forwarded-For
	// header to deduce the original client IP from the trust-
	// worthy parts of that header.
	// Otherwise, simply return the direct client IP
	fmt.Printf("ClientIP: %s\n", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func handleConvert(c *gin.Context) {
	fileName := c.Request.FormValue("name")
	file, _ := c.FormFile("file")
	f, _ := os.CreateTemp("", "*-"+fileName)

	log.Println(fileName)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, f.Name())
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", f.Name()))
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", handleRoot)
	router.POST("/convert", handleConvert)

	router.Run()
}
