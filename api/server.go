package api

import (
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
}

func ListenAndServe(addr string) {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Routes
	router.GET("/", Index)
	router.POST("/", BadVerb)
	router.GET("/healthcheck", Index)
	router.POST("/healthcheck", BadVerb)
	router.GET("/request", BadVerb)
	router.POST("/request", RequestHandler)

	if addr == "" {
		addr = ":80"
	}

	pm := endless.NewServer(addr, router)
	pm.BeforeBegin = func(add string) {
		log.Printf("Server is running at %s", addr)
		log.Printf("Server is running pid is %d", syscall.Getpid())
	}

	pm.ListenAndServe()
}
