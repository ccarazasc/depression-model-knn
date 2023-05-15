package models

type RowData struct {
	Variance      string `json:"variance"`
	TotalDistance string `json:"total_distance"`
	MovementTime  string `json:"movement_time"`
	AverageSpeed  string `json:"average_speed"`
	TestScore     string `json:"test_score"`
	Result        string `json:"result"`
	Terms         map[int]float64
	ResultId      int
}
