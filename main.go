package main

import (
	"log"

	"github.com/Ansh5461/HospitalManagement/database"
	"github.com/Ansh5461/HospitalManagement/handler"
	"github.com/Ansh5461/HospitalManagement/router"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default() //get the default engine for further customizatikon
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	} //get the customized engine

	database.Setup()
	router.SetupRouter(engine, api)
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
