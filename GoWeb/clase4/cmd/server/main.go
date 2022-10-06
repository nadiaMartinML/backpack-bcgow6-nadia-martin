package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase4/cmd/server/handler"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase4/internal/users"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase4/pkg/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)

	u := handler.NewUser(service)

	router := gin.Default()

	rg := router.Group("/users")
	rg.POST("/", u.SaveUser)
	rg.GET("/", u.GetAllUsers)
	rg.PUT("/:id", u.UpdateUser)
	rg.DELETE("/:id", u.DeleteUser)
	rg.PATCH("/:id", u.UpdatePatch)
	router.Run()
}
