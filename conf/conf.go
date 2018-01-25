package conf

import (
	"path/filepath"
	"fmt"
	"os"
	"github.com/astaxie/beego/config"

	"sync"
)
func getExecDir() string {
	execPath, _ := os.Executable()
	return filepath.Dir(execPath)
}


var (
	cfg  map[string]string
	once sync.Once
)

func Config() map[string]string {
	once.Do(func() {
		configPath := filepath.Join(getExecDir(), "conf/app.conf")
		fmt.Println(configPath)
		initConfig, err := config.NewConfig("ini", configPath)
		if err != nil {
			fmt.Printf("error info %s", err)
		}
		cfg, _ = initConfig.GetSection("default")
	})
	return cfg
}