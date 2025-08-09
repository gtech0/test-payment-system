package main

import (
	"log"
	"test-payment-system/database"
	_ "test-payment-system/docs"
	"test-payment-system/enviroment"
	"test-payment-system/router"
)

func init() {
	enviroment.Load()
	database.Connect()
	database.Sync()
}

// @title			Test API
// @version         1
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001
// @BasePath  /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := router.NewRouter()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
