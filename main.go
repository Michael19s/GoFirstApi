package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Michael19s/go_first_api/config"
	"github.com/Michael19s/go_first_api/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("ECHO_PORT")
	if port == "" {
		port = "8080"
	}

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	config.SetClient(client)

	//port = "1323"
	e := echo.New()
	//e.GET("/", controller.UserGetAll)
	router.RegisterRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
