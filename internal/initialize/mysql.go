package initialize

import (
	"gin-api/internal/global"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func initGorm(dsn string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" { // 设置为debug
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string 字段类型默认最大长度
		DisableDatetimePrecision:  true, // 禁止datetime精度, mysql 5.6之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引， 就要把索引先删掉再重建 mysql5.7之前不支持
		DontSupportRenameColumn:   true, // 用change重命名列 mysql8之前不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10) // 设置最大打开的连接数
	sqlDB.SetMaxIdleConns(5)  // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	global.AIG_DB = db
}
