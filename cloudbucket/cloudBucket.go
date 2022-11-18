package cloudbucket

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		bucket := "mmt-app"

		var err error

		ctx := appengine.NewContext(c.Request)

		storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json")) //your credentials file
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		f, uploadedFile, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
				"file":    uploadedFile,
			})
			return
		}

		defer f.Close()

		sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

		if _, err := io.Copy(sw, f); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		if err := sw.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"Error":   true,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "file uploaded successfully",
			"pathname": u.EscapedPath(),
		})
	}
}

//function get file from google storage

func GetFileFromGoogleStorage() gin.HandlerFunc {
	return func(c *gin.Context) {

		bucket := "mmt-app"
		object := c.Params.ByName("nameObject")

		ctx := appengine.NewContext(c.Request)

		storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json")) //your credentials file
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		u, err := storageClient.Bucket(bucket).Object(object).NewReader(ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		f, err := os.Create(object)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating file",
				"error":   true,
			})
			return
		}

		if _, err := io.Copy(f, u); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "error copy file: " + object,
				"error":   true,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "file downloaded successfully",
			"destination:": "C:\\Users\\thanh\\OneDrive\\Máy tính\\qua-mon-mmt\\" + f.Name(),
		})

	}
}

func DeleteFileFromGoogleStorage() gin.HandlerFunc {
	return func(c *gin.Context) {

		bucket := "mmt-app"
		object := c.Params.ByName("nameObject")

		ctx := appengine.NewContext(c.Request)

		storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json")) //your credentials file
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		u, err := storageClient.Bucket(bucket).Object(object).NewReader(ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		if err := u.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		if err := storageClient.Bucket(bucket).Object(object).Delete(ctx); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "file" + object + " deleted successfully",
		})

	}
}