package controllers

//importamos gin que es un mini-framework que nos da lo necesario para realizar una API REST
import (
	"../kmeans"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func RealizarClustering(c *gin.Context) {

	var m1 = make(map[string]float64)
	var m2 = make(map[float64]string)

	////////////////////////////////////////////////////////////////////
	m1["ADULTO MAYOR"] = 1
	m1["ADULTO"] = 2
	m1["JOVEN"] = 3
	m1["NIÑO"] = 4
	m1["ADOLESCENTE"] = 5

	m1["60 a 69 años"] = 6
	m1["50 a 59 años"] = 7
	m1["70 a 79 años"] = 8
	m1["80 a 89 años"] = 9
	m1["30 a 39 años"] = 10
	m1["40 a 49 años"] = 11
	m1["20 a 29 años"] = 12
	m1["90 años a más"] = 13
	m1["00 a 09 años"] = 14
	m1["10 a 19 años"] = 15

	m1["Masculino"] = 16
	m1["Femenino"] = 17

	m1["SAN MARTIN"] = 18
	m1["LA LIBERTAD"] = 19
	m1["LIMA"] = 20
	m1["LAMBAYEQUE"] = 21
	m1["CALLAO"] = 22
	m1["ANCASH"] = 23
	m1["LORETO"] = 24
	m1["PIURA"] = 25
	m1["ICA"] = 26
	m1["UCAYALI"] = 27
	m1["AMAZONAS"] = 28
	m1["JUNIN"] = 29
	m1["TUMBES"] = 30
	m1["MADRE DE DIOS"] = 31
	m1["HUANCAVELICA"] = 32
	m1["HUANUCO"] = 33
	m1["AREQUIPA"] = 34
	m1["PUNO"] = 35
	m1["MOQUEGUA"] = 36
	m1["AYACUCHO"] = 37
	m1["CUSCO"] =38
	m1["TACNA"] = 39
	m1["APURIMAC"] = 40
	m1["CAJAMARCA"] = 41
	m1["PASCO"] = 42

	m1["MINSA"] = 43
	m1["ESSALUD"] = 44
	m1["PNP/FF.AA"] = 45
	m1["DOMICILIO/ALOJAMIENTO"] = 46
	m1["INPE"] = 47
	m1["CLÍNICA PRIVADA"] = 48



	////////////////////////////////////////////////////////
	m2[1]="ADULTO MAYOR"
	m2[2]="ADULTO"
	m2[3]="JOVEN"
	m2[4]="NIÑO"
	m2[5]="ADOLESCENTE"
	m2[6]="60 a 69 años"
	m2[7]="50 a 59 años"
	m2[8]="70 a 79 años"
	m2[9]="80 a 89 años"
	m2[10]="30 a 39 años"
	m2[11]="40 a 49 años"
	m2[12]="20 a 29 años"
	m2[13]="90 años a más"
	m2[14]="00 a 09 años"
	m2[15]="10 a 19 años"
	m2[16]="Masculino"
	m2[17]="Femenino"
	m2[18]="SAN MARTIN"
	m2[19]="LA LIBERTAD"
	m2[20]="LIMA"
	m2[21]="LAMBAYEQUE"
	m2[22]="CALLAO"
	m2[23]="ANCASH"
	m2[24] = "LORETO"
	m2[25] = "PIURA"
	m2[26] = "ICA"
	m2[27] = "UCAYALI"
	m2[28] = "AMAZONAS"
	m2[29] = "JUNIN"
	m2[30] = "TUMBES"
	m2[31] = "MADRE DE DIOS"
	m2[32] = "HUANCAVELICA"
	m2[33] = "HUANUCO"
	m2[34] = "AREQUIPA"
	m2[35] = "PUNO"
	m2[36] = "MOQUEGUA"
	m2[37] = "AYACUCHO"
	m2[38] = "CUSCO"
	m2[39] = "TACNA"
	m2[40] = "APURIMAC"
	m2[41] = "CAJAMARCA"
	m2[42] = "PASCO"
	m2[43] = "MINSA"
	m2[44] = "ESSALUD"
	m2[45] = "PNP/FF.AA"
	m2[46] = "DOMICILIO/ALOJAMIENTO"
	m2[47] = "INPE"
	m2[48] = "CLÍNICA PRIVADA"



	//////////////////////////////////////////////////////////
	data := []kmeans.Punto{}

	var deaths []models.Death
	models.DB.Find(&deaths)

	for _, death := range deaths {
		data = append(data, kmeans.Punto{[]float64{
			m1[death.CATEGORIA],m1[death.DEPARTAMENTO],m1[death.EDAD], m1[death.ETAPA_DE_VIDA], m1[death.SEXO]}})
	}
	param := c.Param("k")
	k, _ := strconv.Atoi(param)
	var clusters = kmeans.KMEANS(data, uint64(k), 5)

	var listRes [][]models.Death

	if len(clusters) > 0 {
		for _, c := range clusters {
			var resultado []models.Death
			for _, p := range c.Puntos {
				var d models.Death
				d.CATEGORIA = m2[p.Entrada[0]]
				d.DEPARTAMENTO = m2[p.Entrada[1]]
				d.EDAD = m2[p.Entrada[2]]
				d.ETAPA_DE_VIDA = m2[p.Entrada[3]]
				d.SEXO = m2[p.Entrada[4]]

				resultado = append(resultado, d)
			}
			listRes = append(listRes, resultado)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": listRes})

}

func GetDeaths(c *gin.Context) {
	var deaths []models.Death
	models.DB.Find(&deaths)
	c.JSON(http.StatusOK, gin.H{"data": deaths})
}
