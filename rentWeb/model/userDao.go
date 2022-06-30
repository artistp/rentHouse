package model

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"rentWeb/utils"
)

func CheckImgCode(uuid, code string) bool {
	conn:=utils.RedisPool.Get()
	defer conn.Close()

	imgCode,err:=redis.String(conn.Do("get",uuid))
	if err != nil {
		log.Fatal(err)
	}

	return imgCode==code
}

func SaveSmsCode(email, code string) error {
	//从redis连接池中获取一个连接
	conn:=utils.RedisPool.Get()
	defer conn.Close()

	_,err:=conn.Do("setex",email+"_code",60*3,code)
	return err
}