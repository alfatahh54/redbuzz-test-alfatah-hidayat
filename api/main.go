package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/alfatahh54/create-transaction/db"
	"github.com/alfatahh54/create-transaction/routes"
	"github.com/gin-gonic/gin"
)

var nextFS embed.FS

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	routes.Router(api)
	distFS, err := fs.Sub(nextFS, "nextjs/dist")
	if err != nil {
		log.Fatal(err)
	}

	// The static Next.js app will be served under `/`.
	http.Handle("/", http.FileServer(http.FS(distFS)))
	port := ":8080"
	db.Database.Migrate()
	getPort := os.Getenv("PORT")
	if getPort != "" {
		port = ":" + getPort
	}
	r.NoRoute(routes.ReverseProxy)
	r.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
