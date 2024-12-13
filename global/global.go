package global

import (
	"gin-api/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	CONFIG              config.Server
	ROUTERS             gin.RoutesInfo
	VP                  *viper.Viper
	DB                  *gorm.DB
	REDIS               *redis.Client
	LOG                 *zap.Logger
	Concurrency_Control = &singleflight.Group{}
)
