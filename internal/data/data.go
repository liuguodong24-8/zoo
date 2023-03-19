package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDB, NewGreeterRepo, NewStudentRepo, NewPledgeRepo, NewAnimalRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	gormDB    *gorm.DB
	ethClient *EthGrpcClient
	c         *conf.Data
}

func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	return db, err
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	client, close, err := NewEthClient("http://1.2.3")
	if err != nil {
		close()
		return nil, func() {}, err
	}

	data := &Data{gormDB: db, ethClient: client, c: c}
	// 启动定时任务
	//go InitTimer(data)
	// go InitTimerAnimal(data)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
