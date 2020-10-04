package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//GetRemainingPower [宿舍号]
//	用于去电费系统获取指定宿舍的最新抄表数据
func GetRemainingPower(room int) (*Power, error) {
	Build := map[string]string{
		"10": "8170",
		"11": "8003",
		"12": "8003",
		"13": "8006",
		"14": "8006",
		"15": "8007",
		"16": "8008",
		"17": "8009",
		"18": "8163",
		"19": "8010",
		"20": "8011",
		"21": "8012",
		"22": "8013",
		"23": "8014",
		"24": "8015",
		"28": "8016",
		"29": "8157",
		"34": "8018",
		"40": "8019",
		"43": "8020",
		"45": "8021",
		"46": "8022",
		"47": "8023",
		"48": "8024",
	}

	//处理输入的房间号
	rooom := strconv.Itoa(room)
	//提取栋数
	build := rooom[0:2]

	//抓取电费信息
	resp, err := http.Get("http://www.gyruibo2.cn/WxSearch/GetRoomInfo?SchID=1&Apartid=" + Build[build] + "&Roomname=" + rooom)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//读取电费信息
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//转为字符串
	pageStr := string(pageBytes)
	//使用选择器选择有效数据
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		return nil, err
	}
	var date [4]string
	var wan uint8
	wan = 0
	dom.Find("div[class=col-xs-5]").Each(func(i int, selection *goquery.Selection) {
		miao, _ := selection.Html()
		// 去除空格
		miao = strings.Replace(miao, "  ", "", -1)
		// 去除换行符
		miao = strings.Replace(miao, "\n", "", -1)
		// 去除Html标签
		miao = strings.Replace(miao, "<label class=\"infolab\">", "", -1)
		miao = strings.Replace(miao, "</label>", "", -1)
		// 去除中文
		miao = strings.Replace(miao, "度", "", -1)
		date[wan] = miao
		wan++
	})

	// 处理得到的数据
	date[1] = strings.Replace(date[1], " ", "", -1)
	date[2] = strings.Replace(date[2], " ", "", -1)
	wwan, _ := strconv.Atoi(date[0])
	date0 := uint16(wwan)
	date1, _ := strconv.ParseFloat(date[1], 64)
	date2, _ := strconv.ParseFloat(date[2], 64)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	the_time, err := time.ParseInLocation("2006/1/02 15:04:05", date[3], loc)
	if err != nil {
		return nil, err
	}
	unix_time := the_time.Unix()
	//这里要转换成标准UTC时间，不然存入数据库后再取出回有问题。
	//遗留问题，有空修复
	local := time.Unix(unix_time, 0)
	date3 := local.UTC()

	//生成struct
	Date := &Power{
		date0,
		date1,
		date2,
		date3,
	}
	return Date, nil
}
