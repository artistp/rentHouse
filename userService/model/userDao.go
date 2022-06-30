package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gomodule/redigo/redis"
	"log"
	"userService/entity"
	"userService/utils"
)

func SaveImgCode(code, uuid string) error{

	conn:=utils.RedisPool.Get()
	defer conn.Close()

	_,err:=conn.Do("setex", uuid,60*5,code)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func CheckImgCode(uuid, code string) bool {
	conn:=utils.RedisPool.Get()
	defer conn.Close()

	imgCode,err:=redis.String(conn.Do("get",uuid))
	if err != nil {
		log.Println(err)
		return false
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

func CheckSmsCode(email,code string)bool{
	conn:=utils.RedisPool.Get()
	defer conn.Close()

	smsCode,err:=redis.String(conn.Do("get",email+"_code"))
	if err != nil {
		log.Println(err)
		return false
	}
	return smsCode==code
}

func SaveRegisterUser(email, password string)error{
	var user entity.User
	user.Email=email
	user.Name=email

	m5:=md5.New()
	m5.Write([]byte(password))
	pwd_hash:=hex.EncodeToString(m5.Sum(nil))  //不实用额外的密钥

	user.Password_hash=pwd_hash

	return utils.MysqlPool.Create(&user).Error
}