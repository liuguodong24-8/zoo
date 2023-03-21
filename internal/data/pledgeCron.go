package data

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron"
	"helloworld/internal/biz"
	"helloworld/internal/conf"
	"helloworld/internal/util"
)

// 定时任务

type pledgeCronRepo struct {
	data *Data
	log  *log.Helper
	conf *conf.Data
}

func (c *pledgeCronRepo) Run() {
	// 一次性扫全表，如果数据量很大，很消耗内存。需要使用 rows.Next
	rows, err := c.data.gormDB.Table("pledges").Where("status = 1").Select("id", "token_id", "address", "created_at", "status", "price", "update_at", "times").Rows()
	if err != nil {
		c.log.Errorf("timer scan fial , err: %v", err)
		return
	}
	// 释放连接
	defer rows.Close()

	param, err := c.data.ethClient.GetPrivate(context.TODO(), c.conf.Eth.ZooPrivate)
	zoo, _ := util.NewZoo(common.HexToAddress(c.conf.Eth.ZooAddr), c.data.ethClient)

	var pledge biz.Pledge
	// 每行单独加载执行业务
	for rows.Next() {
		c.data.gormDB.ScanRows(rows, &pledge)
		// 状态为正在质押，且时间超过7天
		//if pledge.Status == 1 && ((pledge.UpdateAt.Unix() + 7*24*3600) < time.Now().Unix()) {
		//	c.data.gormDB.Table("pledges").Where("id", pledge.ID).Update("times", pledge.Times+1)
		//}
		hash, err := zoo.AddEgg(param, common.HexToAddress(pledge.Address))
		if err != nil {
			c.log.Error(err)
		}
		c.log.Info("Hash:%s", hash)
		//fmt.Println(c.data.ethClient.BlockNumber(context.TODO()))
	}
}

func InitTimer(data *Data) {
	Conrs := cron.New() // 定时任务
	Conrs.Start()
	// 每五秒： */5 * * * * ?
	//  每隔1分钟执行一次："0 */1 * * * ?"
	// 每天23点执行一次："0 0 23 * * ?"
	// 每天凌晨1点执行一次："0 0 1 * * ?"
	// 每月1号凌晨1点执行一次："0 0 1 1 * ?"
	// 在26分、29分、33分执行一次："0 26,29,33 * * * ?"
	// 每天的0点、13点、18点、21点都执行一次："0 0 0,13,18,21 * * ?"

	Conrs.AddJob("*/40 * * * * ?", &pledgeCronRepo{
		data: data,
		log:  nil,
		conf: data.c,
	})
}
