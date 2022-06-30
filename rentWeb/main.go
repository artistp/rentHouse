package main

import (
	"github.com/gin-gonic/gin"
	"rentWeb/controller"
)

func main()  {
	router:=gin.Default()
	router.Static("/home","view")

	rg1:=router.Group("/api/v1.0")
	{
		rg1.GET("/session", controller.GetSession)
		rg1.GET("/imagecode/:uuid",controller.GetImageID)
		rg1.GET("/smscode/:email",controller.GetSmsID)
		rg1.POST("/users", controller.PostRegister)
		rg1.GET("/areas", controller.GetArea)
	}
	router.Run()
}

//func init() {
//	utils.InitRedisPool()
//	utils.InitMysqlPool()
//}
