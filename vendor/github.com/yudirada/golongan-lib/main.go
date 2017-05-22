package lib

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yudirada/golongan-lib/p"
)

func Run() {
	port := os.Getenv("PORT")
	//port = "4536"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.New()
	r.Use(gin.Logger())
	baseURL := "vendor/github.com/yudirada/golongan-lib/"
	r.LoadHTMLGlob(baseURL + "t/*.html")
	//r.Static("/css", "css")
	//r.Static("/script", "script")

	p.RenderIndex(r)

	r.Run(":" + port)
}
