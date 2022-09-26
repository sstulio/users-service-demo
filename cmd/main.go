package main

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/gin-gonic/gin"
	"github.com/sstulio/users-service-demo/internal/app"
	"github.com/sstulio/users-service-demo/internal/pkg/config"
	"github.com/sstulio/users-service-demo/internal/pkg/database"
)

var environment config.Environment

func init() {
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := database.InitDatabase(environment.DatabaseDNS)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/api/users/", app.GetUsers(db))
	r.POST("/api/users/", app.CreateUser(db))

	if err := r.Run(); err != nil {
		log.Fatalf("could not initiate web server %v", err)
	}

}
