package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kylerwsm/bounce/pkg/services"
)

func init() {
	if env, ok := os.LookupEnv("ENV"); ok {
		if isProduction := strings.TrimSpace(env) == "production"; isProduction {
			setupProdEnv()
			return
		}
	}
	setupDevEnv()
}

func setupDevEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupProdEnv() {
	gin.SetMode(gin.ReleaseMode)
}

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
