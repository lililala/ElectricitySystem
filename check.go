package main

import (
	"log"
	"os"
)

//Exists 检查文件/目录是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//dropErr 错误处理
func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

//StartCheck 用于检查各类程序需要的东西是否存在与启动他们
func StartCheck() {
	//这里的负责检查电费系统
	//	检查数据库是否存在
	if !Exists("./" + ElectricitySystem) {
		log.Println("数据库文件[" + ElectricitySystem + "]不存在！正在进行自动创建。")
		InitializeDB()
	}
}
