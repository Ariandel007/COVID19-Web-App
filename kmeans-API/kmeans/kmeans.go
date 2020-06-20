package kmeans

import (
	"math"
	"math/rand"
	"sync"
)

type Punto struct {
	Entrada []float64
}

//Distancia euclidiana
func (punto1 Punto) distanciaAOtroPunto(punto2 Punto) float64 {
	suma := float64(0)

	var wg sync.WaitGroup
	wg.Add(len(punto1.Entrada))

	for e := 0; e < len(punto1.Entrada); e++ {
		go func(e int) {
			defer wg.Done()
			suma += math.Pow(punto1.Entrada[e]-punto2.Entrada[e], 2)
		}(e)
	}

	wg.Wait()

	return math.Sqrt(suma)
}

// centro: Punto que es seleccionado como uno de los k centroides
// puntos: Puntos que pernetecen a su cluster
type Centroide struct {
	Centro Punto
	Puntos []Punto
}

func (centroide1 *Centroide) reCentrar() float64 {
	// crear un nuevo centroiide que sera un punto, ejemplo: float64{1.0,3.0,5.0,2.0, 5.0}
	nuevoCentroide := make([]float64, len(centroide1.Centro.Entrada))

	for _, punto := range centroide1.Puntos {
		for i := 0; i < len(nuevoCentroide); i++ {
			//ahora llenaremos cada valor (x,y,z,..n) del nuevoCentroide con la suma de cada valor (x,y,z,..n) de su punto
			nuevoCentroide[i] += punto.Entrada[i]
		}
	}

	// calcular el promedio(media) de cada nuevo centroide x/=y -> x = x/y
	for j := 0; j < len(nuevoCentroide); j++ {
		nuevoCentroide[j] /= float64(len(centroide1.Puntos))
	}

	// guardamos el anterior centro del centroide en anteriorCentro
	anteriorCentro := centroide1.Centro
	// asignar nuevo centro al atributo centro
	centroide1.Centro = Punto{nuevoCentroide}
	// devolver la distancia del centro anterior y el actual centro
	return anteriorCentro.distanciaAOtroPunto(centroide1.Centro)
}

// data: los puntos que se quieren agrupar
// k: numero de clusters
// threshold: el valor de variacipn maximo en el que se espera que se reacomode cada centroide respecto a su anterior iteracion
func KMEANS(data []Punto, k uint64, threshold float64) (centroides []Centroide) {
	//se crean k centroides, estos son k puntos seleccionados aleatoriamente
	for i := uint64(0); i < k; i++ {
		centroides = append(centroides, Centroide{Centro: data[rand.Intn(len(data))]})
	}

	convergido := false
	for !convergido {
		///////////////////////////////////////////////////////////////
		var wg sync.WaitGroup
		wg.Add(len(data))
		/////////////////////////////////////////////////////////////
		for i := range data {
			//establecemos la distancia minima como la mayor posible
			distanciaMinima := math.MaxFloat64
			z := 0
			go func(i int) {
				defer wg.Done()
				//j: indice del centroide de la iteracion
				for j, centroide := range centroides {
					//calcular la distanica de un punto a un centroide
					distancia := data[i].distanciaAOtroPunto(centroide.Centro)
					//si la distancia del punto al centroide es menor que la distancia minima, entonces esta se vuelve la distancia minima
					if distancia < distanciaMinima {
						distanciaMinima = distancia
						//z se vuelve el indice que corresponde al Centroide mas cercano al punto de la iteracion correspondiente
						z = j
					}
				}
				// añadimos el punto mas cercano al array puntos del centroide
				centroides[z].Puntos = append(centroides[z].Puntos, data[i])
			}(i)
		}
		wg.Wait()

		// asignamos el menor valor posible a delta maximo
		thresholdMaximo := -math.MaxFloat64

		for i := range centroides {
			// el movimiento es la distancia del dentro recalculado con el anterior
			movimiento := centroides[i].reCentrar()
			//comparando la distanica con el delta maximo
			if movimiento > thresholdMaximo {
				thresholdMaximo = movimiento
			}
		}

		//se compara threshold con el delta maximo, en caso de que threshold sea mayor o igual que thresholdMaximo quiere decir que
		// ya hubo poca variacion en el desplazamiento de los centroides, lo cual quiere decir que ya esta convergido
		if threshold >= thresholdMaximo {
			convergido = true
		} else {
			//eliminar los puntos dentro de cada cluster
			for i := range centroides {
				centroides[i].Puntos = nil
			}
		}
	}

	return centroides
}
