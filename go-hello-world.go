package main

import (
	//remove the following line to not have your deployment tracker
	"net/http"
	"os"

	"github.com/IBM-Bluemix/cf_deployment_tracker_client_go"
	"github.com/gin-gonic/gin"
)

func main() {
	//remove the following line to not have your deployment tracker
	cf_deployment_tracker.Track()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})

	router.GET("/loaderio-a7847355da506f3e116029f3839e6e41", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "loaderio-a7847355da506f3e116029f3839e6e41",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
