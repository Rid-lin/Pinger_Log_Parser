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
	// s.DelHeader()
	// s.AddHeader()
	http.Redirect(w, r, "/", 302)
}

func (s *Configuration) editHandler(w http.ResponseWriter, r *http.Request) {
	writetmpl, err := template.ParseFiles("template/write.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	IP := r.FormValue("IP")
	serverElm, ok := s.ServersList[IP]
	if !ok {
		http.NotFound(w, r)
	}

	err = writetmpl.ExecuteTemplate(w, "write", serverElm)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func (s *Configuration) writeHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Configuration) addserverHandler(w http.ResponseWriter, r *http.Request) {
	var serverElm ServersAttr
	serverElm.IP = r.FormValue("IP")
	serverElm.Note = r.FormValue("Note")
	serverElm.SiteID = r.FormValue("SiteID")
	runOncePing(serverElm.IP)
	s.ServersList[serverElm.IP] = serverElm
	tos.fillShapku(servers.ServersList)
	http.Redirect(w, r, "/", 302)
}

func (s *Configuration) deleteHandler(w http.ResponseWriter, r *http.Request) {
	IP := r.FormValue("IP")
	if IP == "" {
		http.NotFound(w, r)
	}
	delete(s.ServersList, IP)
	http.Redirect(w, r, "/", 302)
}

func (s *Configuration) checkNowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Проверяю список адресов... \n")
	runPinger()
	http.Redirect(w, r, "/", 302)
}

//ReLoadDefaultConfigHandler Reload Config servers
func (s *Configuration) ReLoadDefaultConfigHandler(w http.ResponseWriter, r *http.Request) {
	conf := getConf("./default_config.json")
	s = &conf
	s.checkNowHandler(w, r)
	fmt.Printf("Ждём минут: %d\n", servers.TimeOutSleep)
	fmt.Println("================================================================================")
}

//ReLoadConfigHandler Reload Config servers
func (s *Configuration) ReLoadConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Читаю конфигурационный файл config.json\n")
	conf := getConf("./default_config.json")
	s = &conf
	http.Redirect(w, r, "/", 302)
}

//SaveConfigHandler Save configuration in file JSON
func (s *Configuration) SaveConfigHandler(w http.ResponseWriter, r *http.Request) {
	err := backupConfig("./config.json")
	if err != nil {
		toLog(s.logLevel, 2, "Rename file failed: ", err)
	}

	saveConf("./config.json", s)
	http.Redirect(w, r, "/", 302)
}
