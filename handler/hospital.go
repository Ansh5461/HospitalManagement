package handler

import (
	"log"
	"net/http"

	"github.com/Ansh5461/HospitalManagement/database"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPatient(in *gin.Context) {
	patient, err := database.GetPatient(h.DB)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, patient)
}

func (h *Handler) GetPatientRoom(in *gin.Context) {
	name := in.Params.ByName("patient_name")
	patient, err := database.GetPatientRoom(h.DB, name)
	if err != nil {
		in.JSON(http.StatusInternalServerError, err)

	}
	in.JSON(http.StatusOK, patient)
}
