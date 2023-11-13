package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"netsFrame/config"
	"netsFrame/dao/sql"
	"netsFrame/logs"
)

type Module struct {
	*redis.Client
	Ctx    context.Context
	sql    *sql.Query
	player *PlayerInfo
}

var Redis = newRedisModel()

func newRedisModel() *Module {
	redisConf := config.GetGlobalObject().Redis
	if redisConf.Address == "" {
		return nil
	}
	DB := &Module{}
	DB.sql = sql.GetSqlQuery()
	DB.player = NewPlayerInfo()
	DB.Client = redis.NewClient(&redis.Options{
		Addr:     redisConf.Address,
		Username: redisConf.Username,
		Password: redisConf.Password,
	})
	if val, err := DB.Ping(context.Background()).Result(); val != "PONG" {
		logs.PrintLogPanic(err)
	}

	return DB
}
