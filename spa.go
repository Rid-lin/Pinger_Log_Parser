package main

// main project main.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAdresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tos.Data)
}

func getAdress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item := tos.Data[params["IP"]]
	json.NewEncoder(w).Encode(item)
	json.NewEncoder(w).Encode(&LineOfStatusTableType{})
}

func createAdress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var adress LineOfStatusTableType
	_ = json.NewDecoder(r.Body).Decode(&adress)
	tos.Data[adress.IP] = adress
	json.NewEncoder(w).Encode(adress)
	json.NewEncoder(w).Encode(tos.Data)
	runOncePing(adress.IP)
	server := servers.Data[adress.IP]
	server.IP = adress.IP
	server.Note = adress.Note
	server.SiteID = adress.SiteID
	servers.Data[adress.IP] = server
}

func updateAdress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var adress LineOfStatusTableType
	_ = json.NewDecoder(r.Body).Decode(&adress)
	adress.IP = params["IP"]
	tos.Data[adress.IP] = adress
	json.NewEncoder(w).Encode(adress)
	json.NewEncoder(w).Encode(tos.Data)
}

func deleteAdress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	delete(tos.Data, params["IP"])
	json.NewEncoder(w).Encode(tos.Data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./assets/index.html")
	if err != nil {
		fmt.Println("Ошибка чтения")
	}
	// tos.clearCache()
	date := time.Now().Format("2006_01_02")
	tos.storageToCache(date)
	tos.checkactualListIP(&servers)

	fmt.Fprintf(w, "%s", file)
}

func runSpa() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.HandleFunc("/", indexHandler).Methods("GET")

	r.HandleFunc("/api/v1/adresses", getAdresses).Methods("GET")
	r.HandleFunc("/api/v1/adress/{IP}", getAdress).Methods("GET")
	r.HandleFunc("/api/v1/adress", createAdress).Methods("POST")
	r.HandleFunc("/api/v1/adress/{IP}", updateAdress).Methods("PUT")
	r.HandleFunc("/api/v1/adress/{IP}", deleteAdress).Methods("DELETE")
	r.HandleFunc("/api/v1/saveconf", servers.SaveConfigHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
