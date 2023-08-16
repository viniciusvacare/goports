package main

import "github.com/gin-gonic/gin"

func main() {
	// inicializa o router utilizando as config default do gin
	r := gin.Default()

	// define uma rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// executa a api
	r.Run()
}
