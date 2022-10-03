package main

import "github.com/gin-gonic/gin"

func saludoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola Nani",
	})
}

// func main() {
// 	router := gin.Default()

// 	router.GET("/saludo", saludoHandler)

// 	router.Run()
// }
