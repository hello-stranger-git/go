package viperutils

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

// viper初始化
func Setup() {
	viper.SetConfigType("YAML")
	data, err := ioutil.ReadFile("config/config.yaml")
	// 读取配置文件内容
	if err != nil {
		log.Fatalf("Read 'Config.yaml' fail: %v\n", err)
	}
	_ = viper.ReadConfig(bytes.NewBuffer(data))
}
