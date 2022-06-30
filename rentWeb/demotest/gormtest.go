package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"rentWeb/utils"
)

type Student struct {
	gorm.Model
	Name string `gorm:"size:50"`
	Age  int
}

func main() {
	//连接数据库---格式：用户名:密码@协议(IP:端口)/数据库名
	utils.InitMysqlPool()

	//创建表
	//fmt.Println(utils.MysqlPool.AutoMigrate(new(Student)).Error)

	//插入数据
	//var stu Student
	//stu.Name="xiaowu"
	//stu.Age=18
	//err:=utils.MysqlPool.Create(&stu).Error
	//fmt.Println(err)

	//查询数据
	/*
		var stu Student
		//utils.MysqlPool.First(&stu)
		//utils.MysqlPool.Select("name,age").First(&stu) //指定需要查询的字段用select
		utils.MysqlPool.Select("name,age").Where("name=?","xiaowang").First(&stu)
		fmt.Println(stu)
	*/

	//更新数据
	//utils.MysqlPool.Model(new(Student)).Where("name=?","xiaowu").Update("name","xiaowang")//更新一列
	//utils.MysqlPool.Model(new(Student)).Where("name=?","xiaowang").Updates(map[string]interface{}{"name":"xiaoli","age":20})
}
