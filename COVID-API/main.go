package main

import (
	"./controllers"
	"./models"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	// conectar a la base de datos
	models.ConnectDatabase()

	r.GET("/clusters/:k", controllers.RealizarClustering)
	r.GET("/data", controllers.GetDeaths)

	r.POST("/prediccion/:k", controllers.RealizarPrediccion)
	//correr el servidor
	r.Run()
}
