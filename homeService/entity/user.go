package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model //用户编号
	Name          string `gorm:"size:32;unique"` //用户名
	Password_hash string `gorm:"size:128"`
	Email         string `gorm:"size:32;unique"`
	Real_name     string `gorm:"size:32"`
	Id_card       string `gorm:"size:20"`
	Avatar_url    string `gorm:"size:256"`
	Houses        []*House
	Orders        []*OrderHouse
}
