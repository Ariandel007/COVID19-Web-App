package controllers

import (
	"net/http"
	"strconv"

	"../kmeans"
	"../knn"
	"../models"
	"../chain"
	"github.com/gin-gonic/gin"
)

var u uint

func RealizarClustering(c *gin.Context) {
	data := []kmeans.Punto{}

	var setAnalisis []models.Analisis
	models.DB.Find(&setAnalisis)

	for _, analisis := range setAnalisis {
		data = append(data, kmeans.Punto{[]float64{
			analisis.Temperatura, float64(analisis.TosSeca),
			float64(analisis.DolorGargante), float64(analisis.DolorCabeza),
			float64(analisis.DificultadRespirar), float64(analisis.PresionPecho), analisis.IncapacidadParaHablar, float64(analisis.Diagnostico)}})
	}
	param := c.Param("k")
	k, _ := strconv.Atoi(param)
	var clusters = kmeans.KMEANS(data, uint64(k), 0.19)

	var listRes [][]models.Analisis

	if len(clusters) > 0 {
		for _, c := range clusters {
			var resultado []models.Analisis
			for _, p := range c.Puntos {
				var d models.Analisis

				d.Temperatura = p.Entrada[0]
				d.TosSeca = uint64(p.Entrada[1])
				d.DolorGargante = uint64(p.Entrada[2])
				d.DolorCabeza = uint64(p.Entrada[3])
				d.DificultadRespirar = uint64(p.Entrada[4])
				d.PresionPecho = uint64(p.Entrada[5])
				d.IncapacidadParaHablar = p.Entrada[6]
				d.Diagnostico = uint64(p.Entrada[7])
				//d.Diagnostico = uint64(i)
				//models.DB.Create(&d)

				resultado = append(resultado, d)
			}
			listRes = append(listRes, resultado)
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": listRes})

}

func GetDeaths(c *gin.Context) {
	fmt.Println(chain.resultadoConsenso)
	var setAnalisis []models.Analisis
	models.DB.Find(&setAnalisis)
	c.JSON(http.StatusOK, gin.H{"data": setAnalisis})
}

type Vecino struct {
	Dato      models.Analisis `json:"neighbor"`
	Distancia float32         `json:"distancia"`
}

func RealizarPrediccion(c *gin.Context) {
	Kparam := c.Param("k")
	kvalue, _ := strconv.Atoi(Kparam)
	//Extrayendo y convirtiendo dataset
	var trainData [][]float32
	var setAnalisis []models.Analisis
	models.DB.Find(&setAnalisis)
	for _, analisis := range setAnalisis {
		trainData = append(trainData, []float32{float32(analisis.Temperatura),
			float32(analisis.TosSeca),
			float32(analisis.DolorGargante),
			float32(analisis.DolorCabeza),
			float32(analisis.DificultadRespirar),
			float32(analisis.PresionPecho),
			float32(analisis.IncapacidadParaHablar),
			float32(analisis.Diagnostico),
			float32(0)})
	}
	//capturando datos de entrada delusuario
	var datosEntrada models.Analisis

	if c.Bind(&datosEntrada) == nil {
		//Entra aqui si se logro convertir body request a tipo Analisis
		prediccion, arrResultado := knn.Predict_classification(trainData, []float32{
			float32(datosEntrada.Temperatura),
			float32(datosEntrada.TosSeca),
			float32(datosEntrada.DolorGargante),
			float32(datosEntrada.DolorCabeza),
			float32(datosEntrada.DificultadRespirar),
			float32(datosEntrada.PresionPecho),
			float32(datosEntrada.IncapacidadParaHablar),
			float32(datosEntrada.Diagnostico),
		}, kvalue)

		var arrVecinos []Vecino
		for _, v := range arrResultado {
			var n Vecino
			n.Dato.Id = 0
			n.Dato.Temperatura = float64(v[0])
			n.Dato.TosSeca = uint64(v[1])
			n.Dato.DolorGargante = uint64(v[2])
			n.Dato.DolorCabeza = uint64(v[3])
			n.Dato.DificultadRespirar = uint64(v[4])
			n.Dato.PresionPecho = uint64(v[5])
			n.Dato.IncapacidadParaHablar = float64(v[6])
			n.Dato.Diagnostico = uint64(v[7])
			n.Distancia = v[8]
			arrVecinos = append(arrVecinos, n)
		}
		u = uint(prediccion)
		//c.JSON(http.StatusOK, gin.H{"data": arrVecinos, "prediccion": prediccion})
		c.JSON(http.StatusOK, gin.H{})
	}
	//c.JSON(http.StatusUnauthorized, gin.H{"status": "Error al convertir data"})

}

func GetPredicion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": u})
}
