package loginfo

import (
	"github.com/op/go-logging"
	"os"
)


var Log = logging.MustGetLogger("example")
//`%{color}%{time:15:04:05.000} %{level:.5s} %{shortfunc} >  %{message}`,
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000}  %{message}`,
)

func InitLog(logFile *os.File,level logging.Level){
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(logFile, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(level, "")
	logging.SetBackend(backend1Leveled, backend2Formatter)
}