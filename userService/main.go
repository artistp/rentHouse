package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consulapi "github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
	userapi "userService/kitex_gen/userapi/userservice"
	"userService/utils"
)

//启动consul: consul agent -dev -client 10.16.65.76
//启动redis:
func main() {
	r, err := consul.NewConsulRegister("10.16.65.76:8500",consul.WithCheck(&consulapi.AgentServiceCheck{
		Interval:                       "7s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "1m",
	}))
	if err != nil {
		log.Fatal(err)
	}

	addr,_:=net.ResolveTCPAddr("tcp",":8802")
	var opts []server.Option
	opts=append(opts, server.WithServiceAddr(addr))
	opts=append(opts,server.WithRegistry(r))
	opts=append(opts,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "user.UserService",
	}))

	svr := userapi.NewServer(new(UserServiceImpl),opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	utils.InitRedisPool()
	utils.InitMysqlPool()
}