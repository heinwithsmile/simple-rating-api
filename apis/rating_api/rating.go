package rating_api

import (
	"config"
	"encoding/json"
	"models"
	"net/http"
)

//AllData Function
func AllData(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		allRatingModel := models.OverallDataModel{
			Db: db,
		}
		alldata, err2 := allRatingModel.AllData()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, alldata)
		}
	}
}

//WeeklyData Function
func WeeklyData(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		allRatingModel := models.OverallDataModel{
			Db: db,
		}
		alldata, err2 := allRatingModel.WeeklyData()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, alldata)
		}
	}
}

//MonthData Function
func MonthData(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		allRatingModel := models.OverallDataModel{
			Db: db,
		}
		alldata, err2 := allRatingModel.MonthData()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, alldata)
		}
	}
}

//YearData Function
func YearData(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		allRatingModel := models.OverallDataModel{
			Db: db,
		}
		alldata, err2 := allRatingModel.YearData()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, alldata)
		}
	}
}

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
func LoadRating(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		ratingModel := models.RatingModel{
			Db: db,
		}
		ratings, err2 := ratingModel.LoadRating()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, ratings)
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
