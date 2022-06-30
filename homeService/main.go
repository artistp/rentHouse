package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consulapi "github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	homeapi "homeService/kitex_gen/homeapi/homeservice"
	"homeService/utils"
	"log"
	"net"
)

func main() {
	r, err := consul.NewConsulRegister("10.16.65.76:8500",consul.WithCheck(&consulapi.AgentServiceCheck{
		Interval:                       "7s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "1m",
	}))
	if err != nil {
		log.Fatal(err)
	}

	addr,_:=net.ResolveTCPAddr("tcp",":8803")
	var opts []server.Option
	opts=append(opts, server.WithServiceAddr(addr))
	opts=append(opts,server.WithRegistry(r))
	opts=append(opts,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "home.HomeService",
	}))

	svr := homeapi.NewServer(new(HomeServiceImpl),opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	utils.InitRedisPool()
	utils.InitMysqlPool()
}