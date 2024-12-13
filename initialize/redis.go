package initialize

import (
	"context"
	"gin-api/config"
	"gin-api/global"
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
		global.LOG.Error("redis connect ping failed, err:", zap.String("name", redisCfg.Host), zap.Error(err))
	}
	global.LOG.Info("redis connect ping response:", zap.String("name", redisCfg.Host), zap.String("pong", pong))
	return client, nil
}

func Redis() {
	redisClient, err := initRedisClient(global.CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.REDIS = redisClient
}
