package helpers

import (
	"go-api/src/models"
	"math"
	"sort"
	"strings"
	"sync"
	"unicode"

	"github.com/reiver/go-porterstemmer"
)

type KNN struct {
	dictionary     map[string]int
	dictionaryLock sync.RWMutex
	postings       map[int][]int
	postingsLock   sync.RWMutex
	classes        map[string]int
	classesLock    sync.RWMutex
	rowsData       []*models.RowData
	documentsLock  sync.RWMutex
}

type Neighbour struct {
	rowData    *models.RowData
	similarity float64
}

type Neighbours []Neighbour

func NewKNN() *KNN {
	return &KNN{dictionary: make(map[string]int), postings: make(map[int][]int), classes: make(map[string]int), rowsData: make([]*models.RowData, 0)}
}
func NewKNN2() *KNN {
	var data models.Data
	data, _ = ReadCSVFromUrl("https://raw.githubusercontent.com/ccarazasc/TA2-Concurrente/main/src/resources/dataset/Reporte_Proyecto_APROBADO.csv")
	knn := NewKNN()
	knn.Training(data.Data[2:])
	return knn
}

var isAlphaNum = func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) }

func (knn *KNN) newRowData(variance string, totalDistance string, movementTime string, averageSpeed string, testScore string, addTerms bool) *models.RowData {
	rowData := &models.RowData{Variance: variance, TotalDistance: totalDistance, MovementTime: movementTime, AverageSpeed: averageSpeed, TestScore: testScore, Result: " ", Terms: make(map[int]float64)}
	terms := make(map[int]int)

	if addTerms {
		knn.dictionaryLock.Lock()
	} else {
		knn.dictionaryLock.RLock()
	}

	for _, term := range strings.FieldsFunc(variance, isAlphaNum) {
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
	for _, term := range strings.FieldsFunc(totalDistance, isAlphaNum) {
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
	for _, term := range strings.FieldsFunc(movementTime, isAlphaNum) {
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
	for _, term := range strings.FieldsFunc(averageSpeed, isAlphaNum) {
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
	for _, term := range strings.FieldsFunc(testScore, isAlphaNum) {
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

func (knn *KNN) Train(variance string, totalDistance string, movementTime string, averageSpeed string, testScore string, result string) {
	rowData := knn.newRowData(variance, totalDistance, movementTime, averageSpeed, testScore, true)

	knn.classesLock.Lock()

	if classId, ok := knn.classes[result]; !ok {
		classId = len(knn.classes)
		knn.classes[result] = classId
	}

	rowData.ResultId = knn.classes[result]

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

func (knn *KNN) Predict(variance string, totalDistance string, movementTime string, averageSpeed string, testScore string, k int) string {
	rowData := knn.newRowData(variance, totalDistance, movementTime, averageSpeed, testScore, false)

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

	resultCount := make(map[int]int)

	// Count classes in k first (or all) neighbours
	for i := 0; i < k && i < len(neighbours); i++ {
		resultCount[neighbours[i].rowData.ResultId] += 1
	}

	resultId := -1
	maxClassCount := 0

	// Find the most popular class
	for id, count := range resultCount {
		if count > maxClassCount {
			resultId = id
			maxClassCount = count
		}
	}

	result := ""

	knn.classesLock.RLock()

	// Find name of the most popular class
	for c, i := range knn.classes {
		if i == resultId {
			result = c
			break
		}
	}

	knn.classesLock.RUnlock()

	return result
}

func (knn *KNN) Predicciones(data []models.RowData) []models.RowData {
	for d, _ := range data {
		data[d].Result = knn.Predict(data[d].Variance, data[d].TotalDistance, data[d].MovementTime, data[d].AverageSpeed, data[d].TestScore, 1)
	}
	println("Prediciendo...")
	return data
}

func (knn *KNN) Training(data []models.RowData) *KNN {
	for d, _ := range data {
		knn.Train(data[d].Variance, data[d].TotalDistance, data[d].MovementTime, data[d].AverageSpeed, data[d].TestScore, data[d].Result)
	}
	println("Entrenado...")
	return knn
}
