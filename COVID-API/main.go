package main

import (
	"./controllers"
	"./models"
	"./chain"
	"github.com/gin-gonic/gin"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
/*CONSENSO VARIABLES*/

/*FIN CONSENSO VARIABLE*/
func init(){
	chain.iniciarCadena()
	chain.resultadoConsenso=1000
}
func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	// conectar a la base de datos
	models.ConnectDatabase()

	r.GET("/clusters/:k", controllers.RealizarClustering)
	r.GET("/data", controllers.GetDeaths)

	r.POST("/prediccion/:k", controllers.RealizarPrediccion)
	r.GET("/pred", controllers.GetPredicion)

	//correr el servidor
	r.Run()
}
