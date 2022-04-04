package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Ansh5461/HospitalManagement/database"
	"github.com/Ansh5461/HospitalManagement/handler"
)

func SetupRouter(router *gin.Engine, api handler.Handler) {

	router.GET("/hospital", api.GetPatient)
	router.POST("/hospital", api.SavePatient)
	router.GET("/hospital/:patient_id", api.DeletePatientByID)
	//router.GET("/hospital/:patient_name", api.GetPatientRoom)
	router.DELETE("/hospital/:patient_id", api.DeletePatientByID)

}

func Router() *gin.Engine {
	r := gin.Default()

	api := handler.Handler{
		DB: database.GetDB(),
	}
	r.GET("/hospital", api.GetPatient)
	r.POST("/hospital", api.SavePatient)
	r.GET("/hospital/:patient_id", api.DeletePatientByID)
	//router.GET("/hospital/:patient_name", api.GetPatientRoom)
	r.DELETE("/hospital/:patient_id", api.DeletePatientByID)

	return r
}

//{
//    "patient_name":"Ansh",
//	"patient_id": "C123",
//	"room_no": 123,
//	"disease":"Dil mei ched"
//}
//
//

//video 11 59:59
