package models

type Analisis struct {
	Id                    uint64  `json:"id"`
	Temperatura           float64 `json:"temperatura"`
	TosSeca               uint64  `json:"tosSeca"`
	DolorGargante         uint64  `json:"dolorGargante"`
	DolorCabeza           uint64  `json:"dolorCabeza"`
	DificultadRespirar    uint64  `json:"dificultadRespirar"`
	PresionPecho          uint64  `json:"presionPecho"`
	IncapacidadParaHablar float64 `json:"incapacidadParaHablar"`
	Diagnostico           uint64  `json:"diagnostico"`
}
