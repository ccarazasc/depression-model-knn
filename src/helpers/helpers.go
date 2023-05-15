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

		if line[5] == "Si" || line[5] == "No" {
			data.Data = append(data.Data, models.RowData{
				Variance:      line[0],
				TotalDistance: line[1],
				MovementTime:  line[2],
				AverageSpeed:  line[3],
				TestScore:     line[4],
				Result:        line[5],
			})
		}
	}
	return data, nil
}
