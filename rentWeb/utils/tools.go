package utils

import (
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"rentWeb/entity"
	"rentWeb/kitex_gen/homeapi/homeservice"
	"rentWeb/kitex_gen/userapi/userservice"
)

func InitDB() {
	MysqlPool.AutoMigrate(new(entity.User),
		new(entity.House),
		new(entity.Area),
		new(entity.OrderHouse),
		new(entity.HouseImage),
		new(entity.Facility))
}

func GetUserServiceClient()(userservice.Client,error){
	r, err := consul.NewConsulResolver("10.16.65.76:8500")
	if err != nil {
		log.Fatal(err)
	}

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))

	c, err := userservice.NewClient("user.UserService", opts...)
	if err != nil {
		log.Fatal(err)
	}
	return c, err
}

func GetHomeServiceClient()(homeservice.Client,error){
	r, err := consul.NewConsulResolver("10.16.65.76:8500")
	if err != nil {
		log.Fatal(err)
	}

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))

	c, err := homeservice.NewClient("home.HomeService", opts...)
	if err != nil {
		log.Fatal(err)
	}
	return c, err
}
