package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Initialize routes

	// run the server
	r.Run()
}
