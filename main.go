package main

import (
	"fmt"
	"go-api/src/api"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("HTTP_PLATFORM_PORT")
	//var port string = "8080"
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/data/depression", api.GetDataDepression)
	apiRouter.HandleFunc("/predict/depression", api.GetPredictDataDepression)
	apiRouter.HandleFunc("/data/anxiety", api.GetDataAnxiety)
	apiRouter.HandleFunc("/predict/anxiety", api.GetPredictDataAnxiety)
	fmt.Printf("Server running  at port %s", port)
	http.ListenAndServe(":"+port, router)
}
