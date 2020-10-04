package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var room int
var timee int

//调用flag，接受启动参数
func init() {
	flag.IntVar(&room, "room", -1, "需要获取电费的宿舍的宿舍号")
	flag.IntVar(&timee, "time", 3600, "获取间隔(秒)")
	log.SetPrefix("[ElectricitySystem]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	//获取参数
	flag.Parse()
	if room == -1 {
		fmt.Println("请输入宿舍号")
	} else {
		log.Println("初始化中...")
		StartCheck()
		go GetPower()
		StartRPC()
	}
}

func GetPower() {
	for {
		miao, err := GetRemainingPower(room)
		if err != nil {
			log.Println(err)
		} else {
			//向数据库中写入数据
			UpdateData(miao, ElectricitySystem)
		}
		time.Sleep(time.Duration(timee) * time.Second)
	}
}
