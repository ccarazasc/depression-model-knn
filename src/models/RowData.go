package models

type RowData struct {
	EstadoId int
	Id  string `json:"id"`
	Titular string `json:"titular"`
	Ruc    string    `json:"ruc"`
	TituloProyecto    string    `json:"titulo_proyecto"`
	UnidadProyecto         string    `json:"unidad_proyecto"`
	Tipo    string    `json:"tipo"`
	Actividad       string    `json:"actividad"`
	FechaInicio         string    `json:"fecha_inicio"`
	Estado         string    `json:"estado"`
	Descripcion         string    `json:"descripcion"`
	Longitud         string    `json:"longitud"`
	Latitud         string    `json:"latitud"`
	Resolucion         string    `json:"resolucion"`
	Label         string    `json:"label"`
	Terms map[int]float64
}
