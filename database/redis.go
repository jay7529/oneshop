package database

import (
	"os"
	"time"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDISHOST") + ":" + os.Getenv("REDISPORT"),
	DB:   0,
})

func Setkey(key string, value string, time time.Duration) {

	var err = client.Set(key, value, time).Err()
	if err != nil {
		panic(err)
	}

}

func Getkey(key string) interface{} {

	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val

}

func Delkey(key string) interface{} {

	val, err := client.Del(key).Result()
	if err != nil {
		panic(err)
	}
	return val

}

func Existskey(key string) bool {

	val, err := client.Exists(key).Result()
	if err != nil {
		panic(err)
	}
	return val == 1

}

func SetHkey(key string, field string, value interface{}) {

	var err = client.HSet(key, field, value).Err()
	if err != nil {
		panic(err)
	}

}

func GetHkey(key string, field string) interface{} {

	val, err := client.HGet(key, field).Result()
	if err != nil {
		panic(err)
	}
	return val

}

func ExistsHkey(key string, field string) bool {

	val, err := client.HExists(key, field).Result()
	if err != nil {
		panic(err)
	}
	return val

}

func DelHkey(key string, field string) {

	var err = client.HDel(key, field).Err()
	if err != nil {
		panic(err)
	}

}
