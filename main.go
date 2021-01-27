package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kylerwsm/bounce/pkg/services"
)

func main() {
	r := gin.Default()
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
	r.Run(":8080")
}
