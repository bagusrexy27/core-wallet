package utils

import (
	"log"
	"time"
)

func ErrorLog(message string, err error, panic bool) {
	log.Printf("[IPROC-BE] [ERROR] [%v] %v because %v", time.Now().Format("2006-01-02 15:04:05"), message, err)

	if panic {
		log.Fatalf("[IPROC-BE] [ERROR] [%v] %v because %v", time.Now().Format("2006-01-02 15:04:05"), message, err)
	}
}

func InfoLog(message string) {
	log.Printf("[IPROC-BE] [INFO] [%v] %v", time.Now().Format("2006-01-02 15:04:05"), message)

}
