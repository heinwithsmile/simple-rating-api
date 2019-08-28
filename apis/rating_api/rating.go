package rating_api

import (
	"encoding/json"

	"config"
	"models"
	"net/http"
)

//GetAllData Function
func GetAllData(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		ratingModel := models.RatingModel{
			Db: db,
		}
		ratings, err2 := ratingModel.GetAllData()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, ratings)
		}
	}
}

//LoadRating Function
func LoadRating(response http.ResponseWriter, request *http.Request){
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest,err.Error())
	}else{
		ratingModel := models.RatingModel{
			Db : db,
		}
		ratings, err2 := ratingModel.LoadRating()
		if err2 != nil{
			respondWithError(response,http.StatusBadRequest,err2.Error())			
		}else{
			respondWithJSON(response, http.StatusOK,ratings)
		}
	}
}

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
