package logger

import "fmt"

func Debug(message string) {
	showLog(message, "DEBUG")
}

func Info(message string) {
	showLog(message, "INFO")
}

func Critical(message string) {
	showLog(message, "CRITICAL")
}

func Error(message string) {
	showLog(message, "ERROR")
}

func showLog(message, messageType string) {
	msg := fmt.Sprintf("LOGGER %s: %s", messageType, message)
	fmt.Println(msg)
}
