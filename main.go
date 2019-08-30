package main

import (
	"apis/department_api"
	"apis/rating_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/alldata", rating_api.AllData).Methods("GET")
	router.HandleFunc("/api/ratings", rating_api.GetAllData).Methods("GET")
	router.HandleFunc("/api/load", rating_api.LoadRating).Methods("GET")
	router.HandleFunc("/api/department/finddepartments", department_api.FindAllDepartment).Methods("GET")
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":5001", handler)
	if err != nil {
		fmt.Println(err)
	}
}
