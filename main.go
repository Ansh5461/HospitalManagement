package main

import (
	"github.com/Ansh5461/HospitalManagement/database"
	"github.com/Ansh5461/HospitalManagement/handler"
	"github.com/Ansh5461/HospitalManagement/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB here in main
	}

	router.SetupRouter()
}
