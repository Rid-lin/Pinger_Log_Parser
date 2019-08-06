package main

import (
	"fmt"
)

//ServersAttr - attribute and status of server
type ServersAttr struct {
	IP     string `json:"IP"`
	Note   string `json:"Note"`
	SiteID string `json:"SiteID"`
}

// ListOfServers type for list of servers for check
type ListOfServers struct {
	FileLog       string
	NumberOfCheck int
	TimeOutSleep  int
	Data          map[string]ServersAttr
}

//Servers - list of servers for check
var servers ListOfServers

// StatusType type for status servers
type StatusType struct {
	Code       string
	Background string
	NumPass    int
	NumFail    int
}

//StatusNowType type for dysplay status last check
type StatusNowType struct {
	Code string
}

//StatusOfHourType type for calculate of check
type StatusOfHourType struct {
	NumPass int
	NumFail int
}

//LineOfStatusTableType type for display of table
type LineOfStatusTableType struct {
	IP           string
	Note         string
	SiteID       string
	StatusNow    StatusNowType
	StatusOfHour [24]StatusOfHourType
}

type tableOfStatusType struct {
	Data map[string]LineOfStatusTableType
}

var tos tableOfStatusType

func main() {
	// intro()

	// Подготавливаю мапу для заливки
	tos.Data = make(map[string]LineOfStatusTableType)

	//Подготавливаю мапу для заливки конфига
	servers.Data = make(map[string]ServersAttr)
	// Загружаю конфиг
	servers = getConf("./config.json")
	// servers.LoadConfig("./config.json")
	fmt.Println(servers)
}
