// main project config.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
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

//LoadConfig - Load list of Servers for check thier http-status
func (s *ListOfServers) LoadConfig(nameFile string) {
	configFile, err := os.Open(nameFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&s)
}

//LoadDefaultConfig - Load list of Servers for check thier http-status on default
func (s *ListOfServers) LoadDefaultConfig() {
	s.LoadConfig("./default_config.json")
}

//SaveConfig Save configuration in file JSON
func (s *ListOfServers) SaveConfig(nameFile string) {
	// s.DelHeader()
	j, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	// Добавляем перенос на новую строку
	// для лучшей читабельности
	j = append(j, "\r\n"...)
	// Open s.json in append-mode.
	f, _ := os.OpenFile(nameFile, os.O_WRONLY|os.O_CREATE, 0600)
	// Append our json to s.json
	if _, err = f.Write(j); err != nil {
		panic(err)
	}
	fmt.Printf("Файл config.json сохранён.\n")
}

func backupConfig(nameFile string) {
	srcFolder := "nameFile"
	destFolder := "nameFile" + ".bak"
	cmd := exec.Command("copy", "/y /f", srcFolder, destFolder)
	err := cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
