package controllers

import (
	"net/http"
	"strconv"

	"../kmeans"
	"../models"
	"github.com/gin-gonic/gin"
)

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
	var setAnalisis []models.Analisis
	models.DB.Find(&setAnalisis)
	c.JSON(http.StatusOK, gin.H{"data": setAnalisis})
}
