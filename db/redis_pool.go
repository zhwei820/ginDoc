package db

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
	"ginDoc/conf"

)

var (
	// 定义常量
	CachePool *redis.Pool
)

func init() {
	CachePool = initRedisPool("CACHE_REDIS")
}

func initRedisPool(redisUrlKey string) *redis.Pool {
	redisUrl := conf.Config()[redisUrlKey]
	RC := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisUrl)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return RC

}

func usage()  {

	rc := CachePool.Get()
	rc.Do("SET", "s", 0)
	n, err := rc.Do("GET", "s")
	if err !=nil{
		fmt.Println(n)
	}

	// 用完后将连接放回连接池
	defer rc.Close()
}