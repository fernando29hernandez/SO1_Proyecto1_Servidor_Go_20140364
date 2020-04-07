package main

import (
	"fmt"	
	"log"
	"net/http"	
	"encoding/json"	
	"math"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)
import "github.com/mackerelio/go-osstat/memory"

type PROCESO struct{
		PID int32
		Usuario string
		Estado string
		Memoria float32
		Nombre string
		Proceso *process.Process
}
type struct_datos struct{
		TotalProcesos int
		TotalEjecucion int
		TotalSuspendidos int
		TotalDetenidos int
		TotalZombie int
		Procesos []PROCESO
	}


const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)


func datosmemoriaHandler(response http.ResponseWriter, request *http.Request) {
	/*
	meminfo := &procmeminfo.MemInfo{}
	meminfo.Update()
	var total, consumida, porcentaje_consumo, megabytes float64
	megabytes = 1024 * 1024
	total = (float64) (meminfo.Total()) / megabytes //Tamaño en MB
	consumida = (float64) (meminfo.Used()) / megabytes //Tamaño en MB
	porcentaje_consumo = ((consumida * 100) / total)
	fmt.Println((*meminfo)["MemTotal"])
	response.Header().Set("Content-Type","application/json")
	response.WriteHeader(http.StatusOK)
	type MEMORIA struct {
		Total float64
		Consumida float64
		Porcentaje float64
	}

	datos := MEMORIA{Total : total, Consumida : consumida, Porcentaje : porcentaje_consumo}
	datos_json , _ := json.Marshal(datos)
	
	response.Write(datos_json)
	*/
	memory, _ := memory.Get()
	//meminfo.Update()
	var total, consumida, porcentaje_consumo, megabytes float64
	megabytes = 1024 * 1024
	total = (float64) (memory.Total) / megabytes //Tamaño en MB
	consumida = (float64) (memory.Used) / megabytes //Tamaño en MB
	porcentaje_consumo = ((consumida * 100) / total)
	//fmt.Println((*meminfo)["MemTotal"])
	response.Header().Set("Content-Type","application/json")
	response.WriteHeader(http.StatusOK)
	type MEMORIA struct {
		Total float64
		Consumida float64
		Porcentaje float64
	}

	datos := MEMORIA{Total : total, Consumida : consumida, Porcentaje : porcentaje_consumo}
	datos_json , _ := json.Marshal(datos)
	
	response.Write(datos_json)
	
}

func datosCPUHandler(response http.ResponseWriter, request *http.Request) {
	vmStat,_ := cpu.Percent(0,false);
	promedio_uso := math.Floor(vmStat[0]*100)/100
	response.Header().Set("Content-Type","application/json")
	response.WriteHeader(http.StatusOK)
	type CPU struct {
		Porcentaje float64
	}
	datos := CPU{Porcentaje : promedio_uso}
	datos_json , _ := json.Marshal(datos)
	response.Write(datos_json)

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	handler := cors.Default().Handler(router)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/datosmemoria", datosmemoriaHandler)
	router.HandleFunc("/datoscpu", datosCPUHandler)
	fmt.Println("Servidor corriendo en http://localhost:80/")
	log.Fatal(http.ListenAndServe(":80", handler))
}