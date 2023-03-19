package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron"
	"helloworld/internal/biz"
	"time"
)

// 定时任务

type animalCronRepo struct {
	data *Data
	log  *log.Helper
}

func (c *animalCronRepo) Run() {
	// 一次性扫全表，如果数据量很大，很消耗内存。需要使用 rows.Next
	rows, err := c.data.gormDB.Table("animals").Where("status = 1").Select("id", "status", "token_id", "age", "start_time").Rows()
	if err != nil {
		c.log.Errorf("timer scan fial , err: %v", err)
		return
	}
	// 释放连接
	defer rows.Close()

	var a biz.Animal
	// 每行单独加载执行业务
	for rows.Next() {
		c.data.gormDB.ScanRows(rows, &a)
		// 寿命到了
		if a.Status == 1 && ((a.Age*24*3600 + a.StartTime.Unix()) < time.Now().Unix()) {
			fmt.Println(a.ID)
			c.data.gormDB.Table("animals").Where("id", a.ID).Update("status", 0)
		}
	}
}

func InitTimerAnimal(data *Data) {
	Conrs := cron.New() // 定时任务
	Conrs.Start()
	// 每五秒： */5 * * * * ?
	//  每隔1分钟执行一次："0 */1 * * * ?"
	// 每天23点执行一次："0 0 23 * * ?"
	// 每天凌晨1点执行一次："0 0 1 * * ?"
	// 每月1号凌晨1点执行一次："0 0 1 1 * ?"
	// 在26分、29分、33分执行一次："0 26,29,33 * * * ?"
	// 每天的0点、13点、18点、21点都执行一次："0 0 0,13,18,21 * * ?"

	Conrs.AddJob("*/5 * * * * ?", &animalCronRepo{
		data: data,
		log:  nil,
	})
}
