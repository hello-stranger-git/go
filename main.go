package main

import (
	"fmt"
	"go_project/routers"
	mysqlutils "go_project/utils/mysqlUtils"
	viperutils "go_project/utils/viperUtils"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viperutils.Setup()

	// 数据库初始化
	mysqlutils.DB.Init()
	routers := routers.Routers()
	fmt.Printf(http.ListenAndServe(viper.GetString("server.addr"), routers).Error())
}
