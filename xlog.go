package xlog

import (
	"fmt"
	"log"
)

func Debug(msg string, args ...interface{}) {
	log.Printf("[DEBUG] %s\n", fmt.Sprintf(msg, args...))
}

func Warn(msg string, args ...interface{}) {
	log.Printf("[WARN] %s", fmt.Sprintf(msg, args...))
}

func Info(msg string, args ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprintf(msg, args...))
}

func Crit(msg string, args ...interface{}) {
	log.Printf("[CRIT] %s", fmt.Sprintf(msg, args...))
}
