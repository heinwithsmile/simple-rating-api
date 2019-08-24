package department_api

import (
	"config"
	"encoding/json"
	"models"
	"net/http"
)

//FindAllDepartment Function
func FindAllDepartment(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		departmentModel := models.DepartmentModel{
			Db: db,
		}
		cities, err2 := departmentModel.FindAllDepartment()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, cities)
		}
	}
}

//respondWithError Function
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

//respondWithJSON function
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
