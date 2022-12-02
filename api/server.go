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

func ListenAndServe(addr string) error {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Routes
	router.GET("/health", HealcheckHandler)
	router.POST("/request", RequestHandler)

	if addr == "" {
		addr = ":2004"
	}

	pm := endless.NewServer(addr, router)
	pm.BeforeBegin = func(add string) {
		log.Printf("Server is running at %s", addr)
		log.Printf("Server is running pid is %d", syscall.Getpid())
	}

	if err := pm.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
