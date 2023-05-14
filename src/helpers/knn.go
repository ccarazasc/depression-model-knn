package helpers

import (
	"github.com/reiver/go-porterstemmer"
	"go-api/src/models"
	"math"
	"sort"
	"strings"
	"sync"
	"unicode"
)

type KNN struct {
	dictionary     map[string]int
	dictionaryLock sync.RWMutex
	postings       map[int][]int
	postingsLock   sync.RWMutex
	classes        map[string]int
	classesLock    sync.RWMutex
	rowsData      []*models.RowData
	documentsLock  sync.RWMutex
}

type Neighbour struct {
	rowData   *models.RowData
	similarity float64
}

type Neighbours []Neighbour

func NewKNN() *KNN {
	return &KNN{dictionary: make(map[string]int), postings: make(map[int][]int), classes: make(map[string]int), rowsData: make([]*models.RowData, 0)}
}
func NewKNN2() *KNN {
	var data models.Data
	data, _ = ReadCSVFromUrl("https://raw.githubusercontent.com/ccarazasc/TA2-Concurrente/main/src/resources/dataset/Reporte_Proyecto_APROBADO.csv")
	knn:=NewKNN()
	knn.Training(data.Data[2:])
	return knn
}

var isAlphaNum = func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) }

func (knn *KNN) newRowData(id string,titular string,ruc string, tituloProyecto string, unidadProyecto string,tipo string,actividad string, fechaInicio string,descripcion string,longitud string,latitud string,resolucion string,label string, addTerms bool) *models.RowData {
	rowData := &models.RowData{Id: id, Titular: titular, Ruc: ruc, TituloProyecto: tituloProyecto, UnidadProyecto: unidadProyecto, Tipo: tipo, Actividad: actividad, FechaInicio: fechaInicio, Estado: " ", Descripcion: descripcion, Longitud: longitud, Latitud: latitud, Resolucion: resolucion, Label: label, Terms: make(map[int]float64)}
	terms := make(map[int]int)

	if addTerms {
		knn.dictionaryLock.Lock()
	} else {
		knn.dictionaryLock.RLock()
	}

	for _, term := range strings.FieldsFunc(id, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(titular, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(ruc, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(tituloProyecto, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(unidadProyecto, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(tipo, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(actividad, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(fechaInicio, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(descripcion, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(longitud, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(latitud, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(resolucion, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(label, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}
	for _, term := range strings.FieldsFunc(resolucion, isAlphaNum) {
		term = porterstemmer.StemString(term)
		termId, ok := knn.dictionary[term]

		if !ok {
			if addTerms {
				termId = len(knn.dictionary)
				knn.dictionary[term] = termId
			} else {
				termId = -1
			}
		}

		terms[termId]++
	}

	if addTerms {
		knn.dictionaryLock.Unlock()
	} else {
		knn.dictionaryLock.RUnlock()
	}

	// Calculate document's magnitude
	rowDataMagnitude := float64(0.0)

	for _, count := range terms {
		rowDataMagnitude += math.Pow(float64(count), 2)
	}

	rowDataMagnitude = math.Sqrt(rowDataMagnitude)

	// Put terms to the document, normalize their counts by document's magnitude
	for term, count := range terms {
		rowData.Terms[term] = float64(count) / rowDataMagnitude
	}

	return rowData
}

func (knn *KNN) Train(id string,titular string,ruc string, tituloProyecto string, unidadProyecto string,tipo string,actividad string, fechaInicio string,descripcion string,longitud string,latitud string,resolucion string,label string, estado string) {
	rowData := knn.newRowData(id,titular ,ruc , tituloProyecto , unidadProyecto ,tipo ,actividad , fechaInicio ,descripcion ,longitud ,latitud ,resolucion ,label , true)


	knn.classesLock.Lock()

	if classId, ok := knn.classes[estado]; !ok {
		classId = len(knn.classes)
		knn.classes[estado] = classId
	}

	rowData.EstadoId = knn.classes[estado]

	knn.classesLock.Unlock()

	knn.documentsLock.Lock()

	knn.rowsData = append(knn.rowsData, rowData)
	rowDataId := len(knn.rowsData) - 1

	knn.documentsLock.Unlock()

	knn.postingsLock.Lock()

	for termId, _ := range rowData.Terms {
		knn.postings[termId] = append(knn.postings[termId], rowDataId)
	}

	knn.postingsLock.Unlock()
}

func (n Neighbours) Len() int           { return len(n) }
func (n Neighbours) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Neighbours) Less(i, j int) bool { return n[i].similarity > n[j].similarity }

func (knn *KNN) Predict(id string,titular string,ruc string, tituloProyecto string, unidadProyecto string,tipo string,actividad string, fechaInicio string,descripcion string,longitud string,latitud string,resolucion string,label string, k int) string {
	rowData := knn.newRowData(id,titular,ruc,tituloProyecto,unidadProyecto,tipo,actividad,fechaInicio,descripcion,longitud,latitud,resolucion,label,false)

	similarities := make(map[int]float64)

	knn.postingsLock.RLock()
	knn.documentsLock.RLock()

	for termId, _ := range rowData.Terms {
		for rowDataId, _ := range knn.postings[termId] {
			similarities[rowDataId] += knn.rowsData[rowDataId].Terms[termId] * rowData.Terms[termId]
		}
	}

	knn.postingsLock.RUnlock()
	knn.documentsLock.RUnlock()

	neighbours := make(Neighbours, 0)

	for rowDataId, similarity := range similarities {
		neighbours = append(neighbours, Neighbour{knn.rowsData[rowDataId], similarity})
	}

	// Sort neighbours by similarity
	sort.Sort(neighbours)

	estadoCount := make(map[int]int)

	// Count classes in k first (or all) neighbours
	for i := 0; i < k && i < len(neighbours); i++ {
		estadoCount[neighbours[i].rowData.EstadoId] += 1
	}

	estadoId := -1
	maxClassCount := 0

	// Find the most popular class
	for id, count := range estadoCount {
		if count > maxClassCount {
			estadoId = id
			maxClassCount = count
		}
	}

	estado := ""

	knn.classesLock.RLock()

	// Find name of the most popular class
	for c, i := range knn.classes {
		if i == estadoId {
			estado = c
			break
		}
	}

	knn.classesLock.RUnlock()

	return estado
}

func (knn *KNN) Predicciones(data []models.RowData) []models.RowData {
	for d,_ := range data{
		data[d].Estado = knn.Predict(data[d].Id,data[d].Titular,data[d].Ruc,data[d].TituloProyecto,data[d].UnidadProyecto,data[d].Tipo,data[d].Actividad,data[d].FechaInicio,data[d].Descripcion,data[d].Longitud,data[d].Latitud,data[d].Resolucion,data[d].Label,1)
	}
	println("Prediciendo...")
	return data
}

func (knn *KNN) Training(data []models.RowData) *KNN {
	for d,_ := range data{
		knn.Train(data[d].Id,data[d].Titular,data[d].Ruc,data[d].TituloProyecto,data[d].UnidadProyecto,data[d].Tipo,data[d].Actividad,data[d].FechaInicio,data[d].Descripcion,data[d].Longitud,data[d].Latitud,data[d].Resolucion,data[d].Label,data[d].Estado)
	}
	println("Entrenado...")
	return knn
}