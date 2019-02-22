// main project pinger.go
package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func runPinger() {
	cmd := exec.Command("Pinger.exe")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	fmt.Printf("%s\n", stdoutStderr)
}

func runOncePing(IP string) {
	cmd := exec.Command("Pinger.exe", "-adress", IP)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	fmt.Printf("%s\n", stdoutStderr)
}

//checkLoop URL periodic
func (s *ListOfServers) checkLoop() {
	for {
		runPinger()
		fmt.Printf("Ждём минут: %d\n", servers.TimeOutSleep)
		fmt.Println("================================================================================")
		time.Sleep(time.Duration(servers.TimeOutSleep) * time.Minute)
	}
}
