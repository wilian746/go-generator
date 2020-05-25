package logger

import "log"

func INFO(message string, data interface{}) {
	if data == nil {
		log.Print(message)
		return
	}
	log.Print(message, data)
}
