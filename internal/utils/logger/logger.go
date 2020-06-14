package logger

import (
	"github.com/sirupsen/logrus" // nolint
	"log"
)

func ERROR(message string, err error) {
	logrus.Error(message, err.Error())
}

func INFO(message string, data interface{}) {
	if data == nil {
		logrus.Info(message)
		return
	}
	logrus.Info(message, data)
}

func PRINT(messages string) {
	log.SetFlags(0)
	log.Println(messages)
}
