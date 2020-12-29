package redis

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var Conn *redis.Client = nil
var Option *redis.Options = nil

func connect() {
	Option = new(redis.Options)
	Option.Addr = dal.Env.Redis.Addr
	Option.Password = dal.Env.Redis.Password
	Option.DB = dal.Env.Redis.Db
	log.Println(Option.DB)
	log.Println(Option.Password)
	log.Println(Option.Addr)
	Conn = redis.NewClient(Option)
}
func Init() {
	stable.Logger.Info("init redis conn")
	connect()
	go func() {
		for true {
			if Conn.Ping().Err() != nil {
				stable.Logger.Info("keep alive:redis ping failed,reconnecting")
				connect()
				time.Sleep(time.Second * 1)
				continue
			}
			stable.Logger.Info("keep alive: redis ping success ")
			time.Sleep(time.Second * 60)
		}
	}()
}
