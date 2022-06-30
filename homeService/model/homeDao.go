package model

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"homeService/entity"
	"homeService/utils"
	"log"
)

func GetAreaFromRedis()([]byte,error){
	conn:=utils.RedisPool.Get()
	areasSlice,err:=redis.Bytes(conn.Do("get","areaData"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return areasSlice, err
}

func GetAreaFromMysql()([]byte,error){
	var areas []entity.Area
	utils.MysqlPool.Find(&areas)
	areasData, _:=json.Marshal(&areas)

	conn:=utils.RedisPool.Get()
	conn.Do("set","areaData",areasData)

	return areasData,nil
}