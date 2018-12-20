package main

import (
	"io"
	"log"
	"os"
)

func outputLog(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logFile := io.MultiWriter(os.Stdout, file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
}

func main() {
	outputLog("output.log")
	_, err := os.Open("nothingFile")
	if err != nil {
		log.Fatalln("Error Open file")
	}
}
