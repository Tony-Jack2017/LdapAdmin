package redis

import (
	"LdapAdmin/config"
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis() {
	var addr = fmt.Sprintf("%s:%s", config.Conf.Redis.Host, config.Conf.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Conf.Redis.Password,
	})
	defer client.Close()
}
