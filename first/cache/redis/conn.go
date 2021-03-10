package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisHost = "127.0.0.1:6379"
	pool      *redis.Pool
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 5 * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println("redis connect failed, err:", err)
				return nil, err
			}
			return c, nil

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			//检测该连接是否可用，检测时间为1分钟
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
