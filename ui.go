// main project ui.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func (status *StatusOfHourType) calculate() {
// 	if status.NumFail != 0 && status.NumPass != 0 {
// 		if status.NumFail > status.NumPass {
// 			status.Background = "bgpalevioletred"
// 			status.Code = "X"
// 		} else if status.NumFail < status.NumPass {
// 			status.Background = "bgyellowgreen"
// 			status.Code = "√"
// 		} else {
// 			status.Background = "bgyellow"
// 			status.Code = "√"
// 		}
// 	} else if status.NumFail == 0 && status.NumPass != 0 {
// 		status.Background = "bggreen"
// 		status.Code = "√"
// 	} else if status.NumFail != 0 && status.NumPass == 0 {
// 		status.Background = "bgred"
// 		status.Code = "X"
// 	} else {
// 		status.Background = "bggrey"
// 		status.Code = "O"
// 	}
// }

func (tos *tableOfStatusType) fillShapku(source map[string]ServersAttr) {
	for IP, value := range source {
		var shapkaLine LineOfStatusTableType
		shapkaLine.IP = value.IP
		shapkaLine.Note = value.Note
		shapkaLine.SiteID = value.SiteID
		tos.ServersList[IP] = shapkaLine
	}
}

func (tos *tableOfStatusType) storageToCache(date string) {
	filename := "./logs/" + date + ".txt"
	f, err := os.Open(filename)
	if err != nil {
		runPinger()
		f, _ = os.Open(filename)
	}
	r := bufio.NewReader(f)
	s, _, _ := r.ReadLine()
	i := 1
	for len(s) > 0 {
		// fmt.Printf("%v ", i)
		tos.parseLogLine(s)
		s, _, _ = r.ReadLine()
		i++
	}
	// tos.fillTable()
}

func (tos *tableOfStatusType) parseLogLine(b []byte) {
	line := string(b)
	hour, err := strconv.ParseInt(line[12:14], 10, 6)
	if err != nil {
		fmt.Printf("%v Ошибка определения часа.\n%v\n", line[12:14], err)
		return
	}
	IP := strings.Trim(line[21:37], " ")
	rttStr := line[37:38]
	status := tos.ServersList[IP]
	if rttStr == "0" {
		status.StatusOfHour[hour].NumFail++
		status.StatusNow.Code = "X"
	} else {
		status.StatusOfHour[hour].NumPass++
		status.StatusNow.Code = "√"
	}
	tos.ServersList[IP] = status
}

// func (tos *tableOfStatusType) fillTable() {
// 	for IP, status := range tos.ServersList {
// 		// if IP == "IP адрес" {
// 		// 	continue
// 		// }
// 		for i := 0; i < 24; i++ {
// 			cellstatus := status.StatusOfHour[i]
// 			// cellstatus.calculate()
// 			status.StatusOfHour[i] = cellstatus
// 		}
// 		tos.ServersList[IP] = status
// 	}
// }

func (tos *tableOfStatusType) clearCache() {
	for IP, line := range tos.ServersList {
		if IP == "IP адрес" {
			continue
		}
		delete(tos.ServersList, IP)
		var status LineOfStatusTableType
		status.IP = line.IP
		status.Note = line.Note
		status.SiteID = line.SiteID
		tos.ServersList[IP] = status
	}

}

// //DelHeader Delete Header from map
// func (tos *tableOfStatusType) DelHeader() {
// 	delete(tos.ServersList, "IP адрес")
// }

//AddHeader Add Header to map
// func (tos *tableOfStatusType) AddHeader() {
// 	var serverElm LineOfStatusTableType
// 	serverElm.IP = "IP адрес"
// 	serverElm.Note = "Описание"
// 	serverElm.SiteID = "SiteID"
// 	serverElm.StatusNow.Code = "Сейчас"
// 	for i := 24; i > 0; i-- {
// 		str := strconv.Itoa(i - 1)
// 		if len(str) == 1 {
// 			serverElm.StatusOfHour[i-1].Code = "0" + str
// 		} else {
// 			serverElm.StatusOfHour[i-1].Code = str
// 		}
// 	}
// 	tos.ServersList["IP адрес"] = serverElm
// }

func (tos *tableOfStatusType) checkactualListIP(servers *Configuration) {
	for IP := range tos.ServersList {
		if _, ok := servers.ServersList[IP]; !ok && IP != "IP адрес" {
			delete(tos.ServersList, IP)
		}
	}
}
