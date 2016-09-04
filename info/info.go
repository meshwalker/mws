package info
/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

func Info(c *gin.Context) {
	action := c.Param("domain")
	url := c.Request.URL.Host

	log.Info("Received Info request because of failure")
	c.JSON(http.StatusNotFound, gin.H{
		"Status": "Meshwalker offline",
		"Host": c.Request.Host,
		"Domain": action,
		"URL": url,
		"Message": "The meshwalker you tried to access is currently not available",
	})
}*/