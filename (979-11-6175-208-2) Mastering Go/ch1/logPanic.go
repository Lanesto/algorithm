package main

import (
	"os"
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT | syslog.LOG_MAIL, os.Args[0])
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
