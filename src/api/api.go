package api

import (
	"encoding/json"
	"go-api/src/helpers"
	"go-api/src/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var knn = helpers.NewKNN2()

func GetData(res http.ResponseWriter, r *http.Request)  {
	var data models.Data
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data, _ = helpers.ReadCSVFromUrl("https://raw.githubusercontent.com/ccarazasc/TA2-Concurrente/main/src/resources/dataset/Reporte_Proyecto_APROBADO.csv")
	res.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.MarshalIndent(data.Data[2:50], "", " ")
	io.WriteString(res, string(jsonBytes))
}

func GetPredictData(res http.ResponseWriter, r *http.Request)  {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("Getting Data from /predict")
	body,_ := ioutil.ReadAll(r.Body)
	log.Println("Body received:")
	var rowsDataJSON []models.RowData
	json.Unmarshal(body,&rowsDataJSON)
	log.Println(len(rowsDataJSON))
	predicciones := knn.Predicciones(rowsDataJSON)
	jsonB, _ := json.Marshal(predicciones)
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonB)
}