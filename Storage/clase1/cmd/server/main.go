package main

import (
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/Storage/clase1/cmd/server/routes"
	dataBase "github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/Storage/clase1/pkg"
)

func main() {
	engine, db := dataBase.ConnectDatase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
