package city_api

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"entities"
	"config"
	"models"
	"net/http"
)
//FindAll Function
func FindAll(response http.ResponseWriter, request *http.Request){
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest,err.Error())
	}else{
		cityModel := models.CityModel{
			Db : db,
		}
		cities, err2 := cityModel.FindAll()
		if err2 != nil{
			respondWithError(response,http.StatusBadRequest,err2.Error())			
		}else{
			respondWithJSON(response, http.StatusOK,cities)
		}
	}
}
//Search Function
func Search(response http.ResponseWriter ,request *http.Request){

	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response,http.StatusBadRequest,err.Error())
	}else{
		cityModel := models.CityModel{
			Db : db,
		}
		products, err2 := cityModel.Search(keyword)
		if err2 != nil{
			respondWithError(response,http.StatusBadRequest,err2.Error())			
		}else{
			respondWithJSON(response, http.StatusOK,products)
		}
	}
}

// //SearchPrices Function
// func SearchPrices(response http.ResponseWriter ,request *http.Request){

// 	vars := mux.Vars(request)
// 	smin := vars["min"]
// 	smax := vars["max"]
// 	min, _ := strconv.ParseFloat(smin,64)
// 	max, _ := strconv.ParseFloat(smax,64)

// 	db, err := config.GetDB()
// 	if err != nil {
// 		respondWithError(response,http.StatusBadRequest,err.Error())
// 	}else{
// 		productModel := models.ProductModel{
// 			Db : db,
// 		}
// 		products, err2 := productModel.SearchPrices(min,max)
// 		if err2 != nil{
// 			respondWithError(response,http.StatusBadRequest,err2.Error())			
// 		}else{
// 			respondWithJSON(response, http.StatusOK,products)
// 		}
// 	}
// }
// //Create Function
func Create(response http.ResponseWriter ,request *http.Request){
	var city entities.City
	err := json.NewDecoder(request.Body).Decode(&city)
	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response,http.StatusBadRequest,err.Error())
	}else{
		cityModel := models.CityModel{
			Db : db,
		}
		 err2 := cityModel.Create(&city)
		if err2 != nil{
			respondWithError(response,http.StatusBadRequest,err2.Error())			
		}else{
			respondWithJSON(response, http.StatusOK,city)
		}
	}
}

// //Update Function
// func Update(response http.ResponseWriter ,request *http.Request){
// 	var product entities.Product
// 	err := json.NewDecoder(request.Body).Decode(&product)
// 	db, err := config.GetDB()
// 	if err != nil {
// 		respondWithError(response,http.StatusBadRequest,err.Error())
// 	}else{
// 		productModel := models.ProductModel{
// 			Db : db,
// 		}
// 		_, err2 := productModel.Update(&product)
// 		if err2 != nil{
// 			respondWithError(response,http.StatusBadRequest,err2.Error())			
// 		}else{
// 			respondWithJSON(response, http.StatusOK,product)
// 		}
// 	}
// }
 //Delete Function
func Delete(response http.ResponseWriter ,request *http.Request) {

	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)

	db, err := config.GetCloudDB()
	if err != nil {
		respondWithError(response,http.StatusBadRequest,err.Error())
	}else{
		cityModel := models.CityModel{
			Db : db,
		}
	RowsAffected, err2 := cityModel.Delete(id)
		if err2 != nil{
			respondWithError(response,http.StatusBadRequest,err2.Error())			
		}else{
			respondWithJSON(response, http.StatusOK,map[string]int64 {
				"RowsAffected": RowsAffected,
			})
		}
	}
}

//respondWithError Function
func respondWithError(w http.ResponseWriter,code int,msg string){
	respondWithJSON(w, code, map[string]string{"error":msg})
}
//respondWithJSON function
func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	response,_ := json.Marshal(payload)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}