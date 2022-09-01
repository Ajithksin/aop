package logger

import (
	"log"
	"os"
	"testing"
)

func TestNewLogService(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	filename, _ := os.Stat(path)
	log.Println(filename)

	want := NewLogService(path, filename.Name())

	got := NewLogService("/workspace/telemetry-aop/logger", "logger")
	if want != got {
		log.Println("Path and file does not exist")
	}
}
