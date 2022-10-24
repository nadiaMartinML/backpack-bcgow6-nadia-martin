package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoTesting/clase2B/clase4/cmd/server/handler"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoTesting/clase2B/clase4/docs"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoTesting/clase2B/clase4/internal/users"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoTesting/clase2B/clase4/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title NaniMM Bootcamp API
// @version 1.0
// @description This API Handle MELI Products by NaniMM.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name NaniMM API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rg := router.Group("/users")
	rg.POST("/", u.SaveUser)
	rg.GET("/", u.GetAllUsers)
	rg.PUT("/:id", u.UpdateUser)
	rg.DELETE("/:id", u.DeleteUser)
	rg.PATCH("/:id", u.UpdatePatch)
	router.Run()
}
