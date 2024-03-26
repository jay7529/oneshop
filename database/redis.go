package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
	DB:   0,
})

func Setkey(id string) {

	var err = client.Set("golang", id, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}

func Getkey(id string) interface{} {

	val, err := client.Get("golang").Result()
	if err != nil {
		panic(err)
	}
	return val

}

func SetHkey(key_name string, field string, value interface{}) {

	var err = client.HSet(key_name, field, value).Err()
	if err != nil {
		panic(err)
	}

}

func GetHkey(key_name string, field string) interface{} {

	val, err := client.HGet(key_name, field).Result()
	if err != nil {
		panic(err)
	}
	return val

}

func ExistsHkey(key_name string, field string) bool {

	val, err := client.HExists(key_name, field).Result()
	if err != nil {
		panic(err)
	}
	return val

}

func DelHkey(key_name string, field string) {

	var err = client.HDel(key_name, field).Err()
	if err != nil {
		panic(err)
	}

}
