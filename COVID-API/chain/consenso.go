package chain
import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"log"
	"strings"
	"strconv"

	"../models"
)

var Addrs []string //registro de direcciones (IPs) de la red
var TrainData [][]float32
var hostServer string
const (
	puerto_registro = 9000
	puerto_notifica = 8000
	puerto_analisis= 8002
	puerto_notifica_opinion=8003
	puerto_notificar_consenso=8004
	//en el server deberia haber un puesto expecial para recibir los concesos,
	//ese host deberia ser hardcodeado
)
const (
	predA=0
	predB=1
)
var direccion_nodo string

var ResultadoConsenso int=-1 // solo server

type SharedData struct{
	Dataset [][]float32
	Addrs []string
	HostServer string
}
type Opinion struct {
	Addr string
	Prediccion   int
}
type DatosEntrada struct{
	PacienteDatos []float32
	Kvalue int
}

var chInfo chan map[string]int
func EscucharOtrasOpiniones(){
	hostname := fmt.Sprintf("%s:%d", direccion_nodo, puerto_notifica_opinion)
	if ln, err := net.Listen("tcp", hostname); err != nil {
		log.Panicln("Can't start listener on", hostname)
	} else {
		defer ln.Close()
		fmt.Println("Escuando otras predicciones en: ", hostname)
		for {
			//HANDLE ANALISIS
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				//go handleConcenso(conn)
			}
		}
	}
}
func sendOpinion(addrRemota string ,opinion Opinion){
	hostRemoto := fmt.Sprintf("%s:%d", addrRemota, puerto_notifica_opinion)
	if conn, err := net.Dial("tcp", hostRemoto); err != nil {
		log.Println("Can't dial", addrRemota)
	} else {
		defer conn.Close()
		fmt.Println("Sending opinion to", addrRemota)

		bytesOpinion, _ := json.Marshal(opinion)          //serializar
		fmt.Fprintf(conn, "%s\n", string(bytesOpinion)) // enviar respuesta mediante la conexion
	}
}
/*
func handleConcenso(conn net.Conn ){
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var opinionRemota Opinion
	if err := dec.Decode(&opinionRemota); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	}


	info := <-chInfo
	info[opinionRemota.Addr] = opinionRemota.Prediccion
	if len(info) == len(Addrs) {
		ca, cb := 0, 0
		for _, pred := range info {
			if pred == predA {
				ca++
			} else {
				cb++
			}
		}
		if ca > cb {
			sendConsenso(predA)
		} else {
			sendConsenso(predB)
		}
		info = map[string]int{}
	}
	go func() { chInfo <- info }()
}
func sendConsenso(consenso int){
	if conn, err := net.Dial("tcp", hostServer); err != nil {
		log.Println("Can't dial server at", hostServer)
	} else {
		defer conn.Close()
		fmt.Println("Sending to server at", hostServer)
		consensoStr:=strconv.Itoa(consenso)
		fmt.Fprintf(conn, "%s\n", consensoStr)
	}
}
*/

func handleNotificacion(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	ip, _ := br.ReadString('\n')
	ip = strings.TrimSpace(ip)
	//matriculamos la ip al arreglo de ips
	Addrs = append(Addrs, ip)
	fmt.Println("Alguien se unio :",Addrs)
}

func EscucharNotificaciones() {
	hostname := fmt.Sprintf("%s:%d", direccion_nodo, puerto_notifica)
	if ln, err := net.Listen("tcp", hostname); err != nil {
		log.Panicln("Can't start listener on", hostname)
	} else {
		defer ln.Close()
		fmt.Println("Listeing new remote user on: ", hostname)
		for {
			//HANDLE ANALISIS
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handleNotificacion(conn)
			}
		}
	}
}

func Notificar(addr, ip string) {
	remotename := fmt.Sprintf("%s:%d", addr, puerto_notifica)
	if conn, err := net.Dial("tcp", remotename); err != nil {
		log.Println("Can't dial", remotename)
	} else {
		defer conn.Close()
		fmt.Println("Sending to", remotename)
		fmt.Fprintf(conn, "%s\n", ip)
	}
}

func NotificaraTodos(ip string) {
	for _, addr := range Addrs {
		Notificar(addr, ip)
	}
}

func handleRegistro(conn net.Conn) {
	defer conn.Close()
	//leer lo que llega por el punto de conexion = ip
	br := bufio.NewReader(conn)
	ip, _ := br.ReadString('\n')
	ip = strings.TrimSpace(ip)
	//mensaje de confirmacion
	var sharedData SharedData
	sharedData.Dataset=TrainData
	sharedData.Addrs=Addrs
	sharedData.HostServer=hostServer

	bytes, _ := json.Marshal(sharedData)          //serializar
	fmt.Fprintf(conn, "%s\n", string(bytes)) // enviar respuesta mediante la conexion
	NotificaraTodos(ip)                      //notifico a la red
	Addrs = append(Addrs, ip)                //actualizar mi arreglo de direcciones
	fmt.Println("Alguien se unio: ",Addrs)
}

func registrarServer() {
	hostname := fmt.Sprintf("%s:%d", direccion_nodo, puerto_registro)
	if ln, err := net.Listen("tcp", hostname); err != nil {
		log.Panicln("Can't start listener on", hostname)
	} else {
		defer ln.Close()
		fmt.Println("Listeing new members on", hostname)
		for {
			//HANDLE ANALISIS
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handleRegistro(conn)
			}
		}
	}
}

func IniciarCadena(){
	direccion_nodo = myIp()
	hostServer=direccion_nodo +":"+ strconv.Itoa(puerto_notificar_consenso)
	fmt.Println("My ip", direccion_nodo)
	chInfo = make(chan map[string]int)
	go func() { chInfo <- map[string]int{} }()
	//rol de servidor
	go registrarServer()
	//establecer una conexión remota
	//rol de cliente

	//rol servidor
	go EscucharNotificaciones()

	go EscucharConsensos()
	EscucharOtrasOpiniones()
}
func EscucharConsensos(){
	hostname := fmt.Sprintf("%s:%d", direccion_nodo, puerto_notificar_consenso)
	if ln, err := net.Listen("tcp", hostname); err != nil {
		log.Panicln("Can't start listener on", hostname)
	} else {
		defer ln.Close()
		fmt.Println("Escuando otras predicciones en: ", hostname)
		for {
			//HANDLE ANALISIS
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handleConsensoDeConsensos(conn)
			}
		}
	}
}
func handleConsensoDeConsensos(conn net.Conn ){
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var opinionRemota Opinion
	if err := dec.Decode(&opinionRemota); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	}
	info := <-chInfo
	fmt.Println("Se recibio consenso de %s -> %s",opinionRemota.Addr,opinionRemota.Prediccion)
	info[opinionRemota.Addr] = opinionRemota.Prediccion
	if len(info) == len(Addrs) {
		ca, cb := 0, 0
		for _, pred := range info {
			if pred == predA {
				ca++
			} else {
				cb++
			}
		}
		if ca > cb {
			ResultadoConsenso=predA
		} else {
			ResultadoConsenso=predB
		}
		info = map[string]int{}
	}
	go func() { chInfo <- info }()
}
func SendDatosEntrada(datosEntrada DatosEntrada){
	fmt.Print("awfasdf")
	for _,addr:=range Addrs{
		fmt.Println(addr)
		sendDatoEntrada(addr,datosEntrada)
	}
}
func BroadcastOpinion(localPrediccion int){
	opinion:=Opinion{direccion_nodo,localPrediccion}
	for _, addr := range Addrs {
		sendOpinion(addr, opinion)
	}
}
func sendDatoEntrada(remoteAddr string,datosEntrada DatosEntrada){
	hostname := fmt.Sprintf("%s:%d", remoteAddr, puerto_analisis)
	if conn, err := net.Dial("tcp", hostname); err != nil {
		log.Println("Can't dial server at", hostname)
	} else {
		defer conn.Close()
		fmt.Println("Broadcasting paciente data to:", hostname)
		bytes,_:=json.Marshal(datosEntrada)
		fmt.Fprintf(conn, "%s\n", string(bytes))
	}
}
//se informa an nuevo miembro : dataset, addr integrantes, ip server
func myIp() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}
func CargarDataset(){
	var setAnalisis []models.Analisis
	models.DB.Find(&setAnalisis)
	for _, analisis := range setAnalisis {
		TrainData = append(TrainData, []float32{float32(analisis.Temperatura),
			float32(analisis.TosSeca),
			float32(analisis.DolorGargante),
			float32(analisis.DolorCabeza),
			float32(analisis.DificultadRespirar),
			float32(analisis.PresionPecho),
			float32(analisis.IncapacidadParaHablar),
			float32(analisis.Diagnostico),
			float32(0)})
	}
}