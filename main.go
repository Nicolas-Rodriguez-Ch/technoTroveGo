package main

import (
	"github.com/gin-gonic/gin"
	"technoTroveServer/src"
)

func main() {
	router := gin.Default()

	src.ConfigeRoutes(router)

	router.Run("localhost:8080")
}
