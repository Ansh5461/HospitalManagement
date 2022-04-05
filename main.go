package main

import (
	"log"

	"github.com/Ansh5461/HospitalManagement/database"
	//"github.com/Ansh5461/HospitalManagement/handler"
	"github.com/Ansh5461/HospitalManagement/router"
)

func main() {

	database.Setup()
	engine := router.Router() //get the default engine for further customizatikon
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
