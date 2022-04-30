package main

import (
	"douban/cmd"
	_ "douban/docs"
	"douban/tool"
	"fmt"
)

// @title           豆瓣作业
// @version         1.1
// @host            localhost:8084
// @description     期末作业
// @termsOfService  https://github.com/winter-homework-in-redrock/douban.git
// @securityDefinitions.apikey CoreAPI
// @name Authorization
// @in header
func main() {
	//err := tool.InitMySQL()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	err := tool.LinkMysql()
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化表格
	err = tool.CreateTables()
	if err != nil {
		fmt.Println(err)
		return
	}
	//err = tool.InitRedis()
	//if err != nil {
	//	return
	//}
	err = tool.InitGrpcClient()
	if err != nil {
		return
	}
	cmd.URL()
}
