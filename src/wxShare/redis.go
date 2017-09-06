package main

import (
	"github.com/garyburd/redigo/redis"
)

func linkToDb() redis.Conn {
	db, _ := redis.Dial("tcp", "localhost:6379")
	return db
}

func getData(key string) string {
	db := linkToDb()
	defer db.Close()
	res, err := redis.String(db.Do("GET", key))
	if err != nil {
		return ""
	}
	return res
}

func setData(key string, value interface{}) {
	db := linkToDb()
	defer db.Close()
	db.Do("SET", key, value)
}
