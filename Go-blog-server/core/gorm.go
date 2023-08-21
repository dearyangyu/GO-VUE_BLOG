package core

import (
	"Go-blog-server/global"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySqlConnect() *gorm.DB {
	if global.Config.MySql.Host == "" {
		global.LOG.Warn("未配置MySql,取消gorm链接.")
		return nil
	}

	dsn := global.Config.MySql.Dsn()
	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev"{
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger.mysqlLogger
	})
	if err != nil {
		global.LOG.Error(fmt.Sprint("[%s] MySql连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpemConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
