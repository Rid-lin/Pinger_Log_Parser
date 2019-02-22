// main project main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func (s *tableOfStatusType) indexHandler(w http.ResponseWriter, r *http.Request) {
	indextmpl, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	date := time.Now().Format("2006_01_02")
	s.clearCache()

	s.storageToCache(date)
	s.checkactualListIP(&servers)
	err = indextmpl.ExecuteTemplate(w, "index", s)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func (s *tableOfStatusType) addHeadersHendler(w http.ResponseWriter, r *http.Request) {
	s.DelHeader()
	s.AddHeader()
	http.Redirect(w, r, "/", 302)
}

func (s *ListOfServers) editHandler(w http.ResponseWriter, r *http.Request) {
	writetmpl, err := template.ParseFiles("template/write.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	IP := r.FormValue("IP")
	serverElm, ok := s.Data[IP]
	if !ok {
		http.NotFound(w, r)
	}

	err = writetmpl.ExecuteTemplate(w, "write", serverElm)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func (s *ListOfServers) writeHandler(w http.ResponseWriter, r *http.Request) {
	writetmpl, err := template.ParseFiles("template/write.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	err = writetmpl.ExecuteTemplate(w, "write", nil)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func (s *ListOfServers) addserverHandler(w http.ResponseWriter, r *http.Request) {
	var serverElm ServersAttr
	serverElm.IP = r.FormValue("IP")
	serverElm.Note = r.FormValue("Note")
	serverElm.SiteID = r.FormValue("SiteID")
	runOncePing(serverElm.IP)
	s.Data[serverElm.IP] = serverElm
	tos.fillShapku(servers.Data)
	http.Redirect(w, r, "/", 302)
}

func (s *ListOfServers) deleteHandler(w http.ResponseWriter, r *http.Request) {
	IP := r.FormValue("IP")
	if IP == "" {
		http.NotFound(w, r)
	}
	delete(s.Data, IP)
	http.Redirect(w, r, "/", 302)
}

func (s *ListOfServers) checkNowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Проверяю список адресов... \n")
	runPinger()
	http.Redirect(w, r, "/", 302)
}

//ReLoadDefaultConfigHandler Reload Config servers
func (s *ListOfServers) ReLoadDefaultConfigHandler(w http.ResponseWriter, r *http.Request) {
	s.LoadDefaultConfig()
	s.checkNowHandler(w, r)
	fmt.Printf("Ждём минут: %d\n", servers.TimeOutSleep)
	fmt.Println("================================================================================")

}

//ReLoadConfigHandler Reload Config servers
func (s *ListOfServers) ReLoadConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Читаю конфигурационный файл config.json\n")
	s.LoadConfig("config.json")
	// s.checkNowHandler(w, r)
	// fmt.Printf("Ждём минут: %d\n", servers.TimeOutSleep)
	// fmt.Println("================================================================================")
	http.Redirect(w, r, "/", 302)
}

//SaveConfigHandler Save configuration in file JSON
func (s *ListOfServers) SaveConfigHandler(w http.ResponseWriter, r *http.Request) {
	backupConfig("config.json")
	s.SaveConfig("config.json")
	http.Redirect(w, r, "/", 302)
}
