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
	AIG_CONFIG              config.Server
	AIG_ROUTERS             gin.RoutesInfo
	AIG_VP                  *viper.Viper
	AIG_DB                  *gorm.DB
	AIG_REDIS               *redis.Client
	AIG_LOG                 *zap.Logger
	GVA_Concurrency_Control = &singleflight.Group{}
)
