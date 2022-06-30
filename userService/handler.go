package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"log"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
	"userService/kitex_gen/userapi"
	"userService/model"
	"userService/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetImgCaptcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetImgCaptcha(ctx context.Context, req *userapi.ImgCaptchaRequest) (resp *userapi.ImgCaptchaResponse, err error) {
	// TODO: Your code here...
	cap := captcha.New()

	//设置字体
	cap.SetFont("config/comic.ttf")

	cap.SetSize(128, 64)

	cap.SetDisturbance(captcha.NORMAL)

	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})

	cap.SetBkgColor(color.RGBA{0, 0, 128, 128}, color.RGBA{255, 255, 10, 255})

	img, str := cap.Create(4, captcha.NUM)

	//存储到redis数据库中，可以考虑使用消息队列
	err = model.SaveImgCode(str, req.Uuid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//序列化图片
	imgbuf, err := json.Marshal(img)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp = &userapi.ImgCaptchaResponse{
		Img: imgbuf,
	}
	return resp, nil
}

// SendSmsCaptcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendSmsCaptcha(ctx context.Context, req *userapi.SmsCaptchaRequest) (resp *userapi.SmsCaptchaResponse, err error) {
	// TODO: Your code here...
	email := req.Email

	imgCode := req.ImgCode
	uuid := req.Uuid

	if model.CheckImgCode(uuid, imgCode) {
		//发送邮箱验证码
		fromUser := "artistp@163.com"
		toUser := email

		auth := smtp.PlainAuth("panghu", fromUser, "KQMESMBNBVUBKCPC", "smtp.163.com")
		to := []string{toUser}

		//生成验证码
		numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		r := len(numeric)
		rand.Seed(time.Now().UnixNano())
		var sb strings.Builder
		for i := 0; i < 6; i++ {
			fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		}
		smsCode := sb.String()

		resp = &userapi.SmsCaptchaResponse{}

		//发送验证码
		str := fmt.Sprintf("From:%s\r\nTo:%s\r\nSubject:verifycode\r\n\r\nyour verifycode is: %s\r\n", fromUser, email, smsCode)
		msg := []byte(str)
		err = smtp.SendMail("smtp.163.com:25", auth, fromUser, to, msg)
		if err != nil {
			resp.Errno = utils.RECODE_SMSERR
			resp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		} else {
			resp.Errno = utils.RECODE_OK
			resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
			//存储邮箱验证码
			err = model.SaveSmsCode(email, smsCode)
			if err != nil {
				resp.Errno = utils.RECODE_DBERR
				resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			}
		}
	} else {
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	return resp, nil
}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *userapi.RegisterUserRequest) (resp *userapi.RegisterUserResponse, err error) {
	// TODO: Your code here...
	resp=&userapi.RegisterUserResponse{}
	if model.CheckSmsCode(req.Email,req.Smscode) {
		err=model.SaveRegisterUser(req.Email,req.Password)
		if err != nil {
			resp.Errno=utils.RECODE_DBERR
			resp.Errmsg=utils.RecodeText(utils.RECODE_DBERR)
		}else{
			resp.Errno=utils.RECODE_OK
			resp.Errmsg=utils.RecodeText(utils.RECODE_OK)
		}
	}else{
		resp.Errno=utils.RECODE_DATAERR
		resp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
	}
	return resp,nil
}
