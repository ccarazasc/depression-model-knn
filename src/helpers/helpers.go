package helpers

import (
	"encoding/csv"
	"go-api/src/models"
	"io"
	"log"
	"net/http"
)

func ReadCSVFromUrl(url string) (models.Data, error) {
	var data models.Data
	csvFile, _ := http.Get(url)
	reader := csv.NewReader(csvFile.Body)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	reader.Comma = ','
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if line[8] == "Aprobado" || line[8] == "Desaprobado" {
			data.Data = append(data.Data, models.RowData{
				Varianza:          line[0],
				DistanciaTotal:    line[1],
				TiempoMovimiento:  line[2],
				PromedioVelocidad: line[3],
				PuntuaciónTest:    line[4],
				Resultado:         line[5],
			})
		}
	}
	return data, nil
}
