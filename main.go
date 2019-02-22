package main

import (
	"fmt"
	"net/http"
)

var tos tableOfStatusType

func main() {
	intro()
	// Загружаю конфиг
	servers.Data = make(map[string]ServersAttr)
	tos.Data = make(map[string]LineOfStatusTableType)
	servers.LoadConfig("./config.json")
	tos.fillShapku(servers.Data)
	tos.AddHeader()

	go runSpa()

	http.HandleFunc("/", tos.indexHandler)
	http.HandleFunc("/getreport", tos.getreportHandler)
	http.HandleFunc("/checknow", servers.checkNowHandler)
	http.HandleFunc("/write", servers.writeHandler)
	http.HandleFunc("/addserver", servers.addserverHandler)
	http.HandleFunc("/edit", servers.editHandler)
	http.HandleFunc("/delete", servers.deleteHandler)
	http.HandleFunc("/loaddefaultconf", servers.ReLoadDefaultConfigHandler)
	http.HandleFunc("/reloadconf", servers.ReLoadConfigHandler)
	http.HandleFunc("/saveconf", servers.SaveConfigHandler)
	http.Handle("/report/", http.StripPrefix("/report/", http.FileServer(http.Dir("./report"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	fmt.Println("Запуск локального WEB-сервера на порту :8089")
	fmt.Println("================================================================================")

	http.ListenAndServe(":8089", nil)
}
