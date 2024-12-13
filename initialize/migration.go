package initialize

import (
	"gin-api/global"
	"gin-api/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Migration 数据库迁移
func Migration(_db *gorm.DB) {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		global.LOG.Error("migrate table failed", zap.Error(err))
		panic(err)
	}
	global.LOG.Info("migrate table success")
}
