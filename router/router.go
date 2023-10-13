package router

import (
	gin "github.com/gin-gonic/gin"
)

func Initializer(){
	// Initialize Router
	router := gin.Default()
	// Initialize Routes
	initializeRoutes(router)
	// Run the server
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}