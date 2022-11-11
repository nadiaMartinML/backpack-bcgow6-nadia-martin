package dataBase

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDatase() (engine *gin.Engine, db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: no se pudo cargar el .env")
	}

	configDataBase := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    os.Getenv("DBNET"),
		Addr:   os.Getenv("DBADDRESS"),
		DBName: os.Getenv("DBNAME"),
	}

	db, err = sql.Open("mysql", configDataBase.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	engine = gin.Default()

	return engine, db
}
