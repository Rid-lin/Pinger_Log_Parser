// main project config.go
package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// //LoadConfig - Load list of Servers for check thier http-status
// func (s *Configuration) LoadConfig(nameFile string) {
// 	configFile, err := os.Open(nameFile)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer configFile.Close()
// 	jsonParser := json.NewDecoder(configFile)
// 	jsonParser.Decode(&s)
// }

// //SaveConfig Save configuration in file JSON
// func (s *Configuration) SaveConfig(nameFile string) {
// 	// s.DelHeader()
// 	j, err := json.MarshalIndent(s, "", "    ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Добавляем перенос на новую строку
// 	// для лучшей читабельности
// 	// Open s.json in append-mode.
// 	f, _ := os.OpenFile(nameFile, os.O_WRONLY|os.O_CREATE, 0600)
// 	// Append our json to s.json
// 	if _, err = f.Write(j); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Файл %v сохранён.\n", nameFile)
// }
