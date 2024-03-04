package main

import (
	middlewares "github.com/jon-d-elliott/address-service/src/middleware"
	"github.com/jon-d-elliott/address-service/src/models"
	"github.com/jon-d-elliott/address-service/src/routes"
	"github.com/jon-d-elliott/address-service/src/utils"
)

func main() {
	utils.LoadEnv()
	router := routes.SetupRoutes()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	middlewares.RegisterMiddlewares(router)
	router.Run(":8113")

}
