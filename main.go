package main

import (
	"apis/city_api"
	"apis/department_api"
	"apis/rating_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/city/findall", city_api.FindAll).Methods("GET")
	router.HandleFunc("/api/city/findrating", rating_api.FindRating).Methods("GET")
	router.HandleFunc("/api/department/finddepartments", department_api.FindAllDepartment).Methods("GET")
	router.HandleFunc("/api/city/search/{keyword}", city_api.Search).Methods("GET")
	// router.HandleFunc("/api/product/search/{min}/{max}", product_api.SearchPrices).Methods("GET")
	router.HandleFunc("/api/city/create", city_api.Create).Methods("POST")
	// router.HandleFunc("/api/product/update", product_api.Update).Methods("PUT")
	router.HandleFunc("/api/city/delete/{id}", city_api.Delete).Methods("DELETE")
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":5001", handler)
	if err != nil {
		fmt.Println(err)
	}
}
