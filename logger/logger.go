package logger

import (
	"log"
	"os"
)

//var errorfaliure = errors.New("faild to setup")

type logger struct {
	file string
	Do   *log.Logger
}

func NewLogService(path string, name string) *logger {
	return createlogger(path, name)
}

func createlogger(path string, name string) *logger {
	//err := os.MkdirAll(path, os.ModePerm)

	file, err := os.OpenFile(path+name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return nil
	}
	return &logger{
		file: name,
		Do:   log.New(file, "Log\t:", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
