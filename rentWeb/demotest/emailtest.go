package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

func main() {
	fromUser:="artistp@163.com"
	toUser:="1164153593@qq.com"

	auth := smtp.PlainAuth("panghu",fromUser,"KQMESMBNBVUBKCPC","smtp.163.com")
	to:=[]string{toUser}

	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < 6; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	num:= sb.String()

	str:=fmt.Sprintf("From:%s\r\nTo:%s\r\nSubject:verifycode\r\n\r\nyour verifycode is: %s\r\n",fromUser, toUser, num)
	msg:=[]byte(str)
	err:=smtp.SendMail("smtp.163.com:25",auth,fromUser,to,msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
