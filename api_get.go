package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var infinit = 1
	for infinit > 0 {
		response, err := http.Get("https://cat-fact.herokuapp.com/facts/random")

		if err != nil {
			writeLog(string(err.Error()), "./logs/error.log")
		}
		defer response.Body.Close()

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			writeLog(string(err.Error()), "./logs/error.log")
		}
		escriureLog := string(responseData)
		writeLog(escriureLog, "./logs/apiDocker.log")

		time.Sleep(4 * time.Second) //Fem consultes a la api cada 3 segons
	}
}

func writeLog(escriureLog string, nomLog string) {
	f, err := os.OpenFile(nomLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //Obrim o creem el fitxer apiDocker.log
	if err != nil {                                                          //Comprovem que no hi ha hagut cap error
		writeLog(string(err.Error()), "./logs/error.log")
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags) //Li diem a la llibreria que volem posar la data de quan escrivim al log
	logger.Println(escriureLog)
}
