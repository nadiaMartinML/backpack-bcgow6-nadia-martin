package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase2B/cmd/server/handler"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase2B/internal/users"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)

	u := handler.NewUser(service)

	router := gin.Default()

	rg := router.Group("/users")
	rg.POST("/", u.SaveUser)
	rg.GET("/", u.GetAllUsers)
	router.Run()
}
