package server

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

	// router.POST("/convert", handleConvert)

	if addr == "" {
		addr = "127.0.0.1:2003"
	}

	server := endless.NewServer(addr, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Server is running at %s", addr)
		log.Printf("Server is running pid is %d", syscall.Getpid())
	}

	server.ListenAndServe()
}
