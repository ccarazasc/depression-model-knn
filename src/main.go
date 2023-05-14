package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-api/src/api"
	"net/http"
)

func main()  {
	var port string = "8080"



	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/data", api.GetData)
	apiRouter.HandleFunc("/predict", api.GetPredictData)
	fmt.Printf("Server running  at port %s", port)
	http.ListenAndServe(":"+port,router)
}
