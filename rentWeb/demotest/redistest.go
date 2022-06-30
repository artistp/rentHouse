package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn,err:=redis.Dial("tcp","10.16.65.76:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reply,err:=conn.Do("set", "name","xiaowu")
	if err != nil {
		log.Fatal(err)
	}

	r,err:=redis.String(reply,err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r,err)

}
