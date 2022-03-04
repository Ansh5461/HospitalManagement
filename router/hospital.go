package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Ansh5461/HospitalManagement/handler"
)

func SetupRouter(router *gin.Engine, api handler.Handler) {

	router.GET("/hospital", api.GetPatient)
	router.POST("/hospital", api.SavePatient)
	router.GET("/hospital/:patient_id", api.DeletePatientByID)
	//router.GET("/hospital/:patient_name", api.GetPatientRoom)
	router.DELETE("/hospital/:patient_id", api.DeletePatientByID)

}
