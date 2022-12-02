package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Healcheck healcheck

func HealcheckHandler(c *gin.Context) {
	if Healcheck.Status == "busy" {
		c.Status(http.StatusServiceUnavailable)
	} else {
		c.Status(http.StatusOK)
	}
}

type healcheck struct {
	Status string
}

func (h *healcheck) SetAvailable() {
	h.Status = "available"
}

func (h *healcheck) SetBusy() {
	h.Status = "busy"
}
