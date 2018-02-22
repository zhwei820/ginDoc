package utils

import (
	clog "github.com/cihub/seelog"

	"os"
	"path/filepath"
	"fmt"
)

func init()  {
	SetSeeLog("conf/seelog.xml")
}

func getExecDir() string {
	execPath, _ := os.Executable()
	return filepath.Dir(execPath)
}

func SetSeeLog(logPath string) bool {
	fmt.Println(logPath)
	logger, err := clog.LoggerFromConfigAsFile(logPath)
	if err != nil {
		fmt.Println("set log config err:", err)
		return false
	}
	if nLogger, ok := logger.(clog.LoggerInterface); ok {
		clog.UseLogger(nLogger)
		clog.Infof("set new logger successfully.")
		return true
	}
	return false
}
