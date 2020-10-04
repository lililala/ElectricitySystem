package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var ElectricitySystem string = "./ElectricitySystem.db"

//UpdateData [需要保存的数据] [数据库路径]
func UpdateData(power *Power, database string) {
	log.Println("开始获取电费...")
	db, err := sql.Open("sqlite3", database+"?charset=utf")
	dropErr(err)
	defer db.Close()

	//判断是否存在相同的数据，避免在暴力循环抓取中uid快速膨胀
	if !JudgeRowExistence(db, power) {
		log.Print(power)
		log.Println("正在写入数据库")
		//更新表latest
		AddDB(db, power, database, "latest")
		//更新表power
		AddDB(db, power, database, "power")
		log.Println("数据库写入完成")
	}
}

//AddDB([需要保存的数据], [数据库路径], [表名])
//	AddDB 插入数据到指定数据库与表
func AddDB(db *sql.DB, power *Power, database string, table string) {
	//插入数据
	stmt, err := db.Prepare("INSERT OR REPLACE INTO " + table + "(Room, Used, Remaining, Date) values(?,?,?,?)")
	dropErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(power.Room, power.Used, power.Remaining, power.Date)
	dropErr(err)
}

//JudgeRowExistence 查询电费系统中是否已经收录该行
func JudgeRowExistence(db *sql.DB, power *Power) bool {
	rows, err := db.Query("select room,date FROM latest where room=" + strconv.Itoa(int(power.Room)))
	dropErr(err)
	defer rows.Close()

	for rows.Next() {
		var room int
		var date time.Time
		err = rows.Scan(&room, &date)
		dropErr(err)
		if date == power.Date {
			log.Println("电费数据未更新")
			return true
		}
		log.Println("电费数据已更新")
		return false
	}
	err = rows.Err()
	log.Println("电费数据已更新")
	return false
}

//InitializeDB 初始化数据库
func InitializeDB() {
	//创建数据库文件
	db, err := sql.Open("sqlite3", ElectricitySystem+"?charset=utf")
	dropErr(err)
	defer db.Close()

	//创建数据库表
	//	创建用于存储历史用电记录的表
	sql_table := `
    CREATE TABLE IF NOT EXISTS power(
        room INTEGER NULL,
        used REAL NULL,
        remaining REAL NULL,
		date DATETIME NULL,
		UNIQUE(room,date)
	);
	`
	db.Exec(sql_table)
	//	创建用于存储已经收录入系统的房间号与最新的用电数据的表
	sql_table = `
	CREATE TABLE IF NOT EXISTS latest(
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
        room INTEGER NULL,
        used REAL NULL,
        remaining REAL NULL,
		date DATETIME NULL,
		UNIQUE(room)
	);
	`
	db.Exec(sql_table)
}
