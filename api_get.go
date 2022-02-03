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

		time.Sleep(4 * time.Second)
	}
}

func writeLog(escriureLog string, nomLog string) {
	f, err := os.OpenFile(nomLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		writeLog(string(err.Error()), "./logs/error.log")
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(escriureLog)
}
