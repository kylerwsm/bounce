package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kylerwsm/bounce/internal/database"
	"github.com/kylerwsm/bounce/pkg/models"
	"github.com/kylerwsm/bounce/pkg/services"
)

func init() {
	db := database.GetDatabase()
	db.AutoMigrate(&models.ShortLink{})
}

func main() {
	r := gin.Default()

	// Redirect route.
	r.GET("/:path", func(c *gin.Context) {
		path, _ := c.Params.Get("path")
		targetURL, err := services.RedirectFrom(path)
		if err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, targetURL)
	})

	// Create route.
	r.POST("/:path", func(c *gin.Context) {
		path, _ := c.Params.Get("path")

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}

		linkStruct := struct {
			OriginalLink string `json:"link"`
		}{}
		if err := json.Unmarshal(data, &linkStruct); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}

		if err := services.CreateLink(path, linkStruct.OriginalLink); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}

		c.Status(http.StatusOK)
	})

	// Delete route.
	r.DELETE("/:path", func(c *gin.Context) {
		path, _ := c.Params.Get("path")
		if err := services.DeleteLink(path); err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}
		c.Status(http.StatusOK)
	})

	r.Run(":8080")
}
