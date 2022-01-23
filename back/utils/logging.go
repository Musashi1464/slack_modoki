package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings() {
	logfile, err := os.OpenFile("./slack_modoki.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer logfile.Close()
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
