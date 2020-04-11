package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"hara-depo-proj/config"
)

func redisConn() *redis.Client {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	client := redis.NewClient(&redis.Options{
		Addr:     baseConfig.Redis.Host + ":" + baseConfig.Redis.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
func SaveData(key string, value string) {
	err := redisConn().Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetData(key string) error {
	err := redisConn().Get(key).Err()
	if err != nil {
		fmt.Println("err")
	}
	return err
}

func GetDataFromRedis(key string) string {
	value, _ := redisConn().Get(key).Result()
	return value
}

func DeleteDataFromRedis(key string) {
	err := redisConn().Del(key).Err()
	if err != nil {
		panic(err)
	}

}
