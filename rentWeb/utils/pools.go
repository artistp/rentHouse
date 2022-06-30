package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"log"
)

var RedisPool redis.Pool
var MysqlPool *gorm.DB

func InitRedisPool(){
	RedisPool=redis.Pool{
		MaxIdle: 20,
		MaxActive: 50,
		MaxConnLifetime: 60*5,
		IdleTimeout: 60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","10.16.65.76:6379")
		},
	}
}

func InitMysqlPool()  {
	var err error
	MysqlPool,err=gorm.Open("mysql","root:wp19970426@tcp(127.0.0.1:3306)/rentHome?parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	MysqlPool.DB().SetMaxIdleConns(10)
	MysqlPool.DB().SetMaxOpenConns(100)
	MysqlPool.SingularTable(true) //不使用复数表名
	//InitDB()
}
