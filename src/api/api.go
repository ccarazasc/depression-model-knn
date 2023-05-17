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

var knnDepression = helpers.NewKNN2()
var knnAnxiety = helpers.NewKNN3()

func GetDataDepression(res http.ResponseWriter, r *http.Request) {
	var data models.Data
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	data, _ = helpers.ReadCSVFromUrl("https://raw.githubusercontent.com/ccarazasc/depression-model-knn/master/src/resources/data-depression.csv")
	res.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.MarshalIndent(data.Data[1:len(data.Data)], "", " ")
	io.WriteString(res, string(jsonBytes))
}

func GetDataAnxiety(res http.ResponseWriter, r *http.Request) {
	var data models.Data
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	data, _ = helpers.ReadCSVFromUrl("https://raw.githubusercontent.com/ccarazasc/depression-model-knn/master/src/resources/data-anxiety.csv")
	res.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.MarshalIndent(data.Data[1:len(data.Data)], "", " ")
	io.WriteString(res, string(jsonBytes))
}

func GetPredictDataDepression(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("Getting Data from /predict")
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("Body received Depression:")
	var rowDataJSON models.RowData
	json.Unmarshal(body, &rowDataJSON)
	log.Println(&rowDataJSON)
	prediccion := knnDepression.Prediccion(rowDataJSON)
	jsonB, _ := json.Marshal(prediccion)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	log.Println(prediccion)
	res.Write(jsonB)
}

func GetPredictDataAnxiety(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("Getting Data from /predict")
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("Body received Anxiety:")
	var rowDataJSON models.RowData
	json.Unmarshal(body, &rowDataJSON)
	log.Println(&rowDataJSON)
	prediccion := knnAnxiety.Prediccion(rowDataJSON)
	jsonB, _ := json.Marshal(prediccion)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	log.Println(prediccion)
	res.Write(jsonB)
}
