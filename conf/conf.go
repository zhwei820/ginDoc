package conf

import (
	"path/filepath"
	"fmt"
	"os"
	"github.com/astaxie/beego/config"
)

func getExecDir() string {
	execPath, _ := os.Executable()
	return filepath.Dir(execPath)
}

var (
	Cfg map[string]string
)

func init() {
	configPath := filepath.Join(getExecDir(), "conf/app.conf")
	fmt.Println(configPath)
	initConfig, err := config.NewConfig("ini", configPath)
	if err != nil {
		fmt.Printf("error info %s", err)
	}
	Cfg, _ = initConfig.GetSection("default")

}
