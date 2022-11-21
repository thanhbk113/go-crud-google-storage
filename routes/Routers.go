package routes

import (
	"my-app/cloudbucket"

	"github.com/gin-gonic/gin"
)

func CloudRoutes(incoming *gin.Engine) {
	incoming.POST("/cloud-storage-bucket", cloudbucket.HandleFileUploadToBucket())
	incoming.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	incoming.GET("/cloud-storage-bucket/:nameObject/:isDirectly", cloudbucket.GetFileFromGoogleStorage())

	incoming.DELETE("/cloud-storage-bucket/:nameObject", cloudbucket.DeleteFileFromGoogleStorage())
}
