package main

import "time"

//宿舍电费数据
//
//	Room	宿舍号
//	Used	该宿舍总计用电量
//	Remaining	当前抄表(剩余电量)电量
//	Date	抄表时间
type Power struct {
	Room      uint16
	Used      float64
	Remaining float64
	Date      time.Time
}
