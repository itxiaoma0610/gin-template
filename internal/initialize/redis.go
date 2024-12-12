package initialize

import (
	"context"
	"gin-api/config"
	"gin-api/internal/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(redisCfg config.Redis) (*redis.Client, error) {
	// 默认使用单列模式
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Pass,
		DB:       0,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.AIG_LOG.Error("redis connect ping failed, err:", zap.String("name", redisCfg.Host), zap.Error(err))
	}
	global.AIG_LOG.Info("redis connect ping response:", zap.String("name", redisCfg.Host), zap.String("pong", pong))
	return client, nil
}

func Redis() {
	redisClient, err := initRedisClient(global.AIG_CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.AIG_REDIS = redisClient
}
