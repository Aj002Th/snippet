package logger

import (
	"log"
	"os"
)

var (
	Logger     *log.Logger
	HttpLogger *log.Logger
)

func Init() {
	Logger = log.New(os.Stdout, "[Logger]", log.LstdFlags|log.Llongfile)
	HttpLogger = log.New(os.Stdout, "[HttpLogger]", log.LstdFlags)
}