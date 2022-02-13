package handler

import (
	"log"
	"net/http"

	"github.com/Ansh5461/HospitalManagement/database"
	"github.com/Ansh5461/HospitalManagement/models"
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
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)

	}
	in.JSON(http.StatusOK, patient)
}

func (h *Handler) GetPatientByID(in *gin.Context) {
	id := in.Params.ByName("patient_id")
	patient, err := database.GetPatientByID(h.DB, id)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)

	}
	in.JSON(http.StatusOK, patient)

}

func (h *Handler) DeletePatientByID(in *gin.Context) {
	id := in.Params.ByName("patient_id")
	err := database.DeletePatientByID(h.DB, id)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, nil)
}

func (h *Handler) SavePatient(in *gin.Context) {
	patient := models.Hospital{}
	err := in.BindJSON(&patient)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}

	err = database.SavePatient(h.DB, &patient)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, &patient)
}
