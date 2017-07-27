package main

import (
	"log"
)

type Logger struct {
}

func (logger *Logger) Info(args ...interface{}) {
	log.Print(append([]interface{}{"[INFO]"}, args...)...)
}
func (logger *Logger) Infof(fmt string, args ...interface{}) {
	log.Printf("[INFO] "+fmt, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	log.Print(append([]interface{}{"[DEBUG]"}, args...)...)
}
func (logger *Logger) Debugf(fmt string, args ...interface{}) {
	log.Printf("[DEBUG] "+fmt, args...)
}
