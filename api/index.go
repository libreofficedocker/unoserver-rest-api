package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	log.Printf("UnoServer is up")
	c.String(http.StatusOK, "UnoServer is up")
	return
}

func BadVerb(c *gin.Context) {
	log.Printf("Method not allowed for this endpoint")
	c.String(http.StatusMethodNotAllowed , "Method not allowed for this endpoint")
	return
}