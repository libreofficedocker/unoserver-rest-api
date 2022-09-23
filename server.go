package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var WorkDir string = "workd"

func init() {
	d, _ := os.MkdirTemp("", "unodata-*")
	WorkDir = d
	log.Printf("Server working directory '%s'", d)
}

func cleanup() {
	log.Printf("Removing directory '%s'", WorkDir)
	os.RemoveAll(WorkDir)
}

func main() {
	defer cleanup()

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", handleRoot)
	router.POST("/convert", handleConvert)

	server := endless.NewServer("localhost:4242", router)
	server.BeforeBegin = func(add string) {
		log.Printf("Server is running pid is %d", syscall.Getpid())
	}
	server.ListenAndServe()
}

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
