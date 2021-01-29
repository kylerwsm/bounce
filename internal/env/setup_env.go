package env

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
