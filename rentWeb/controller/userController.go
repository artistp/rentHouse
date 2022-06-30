package controller

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/gin-gonic/gin"
	"image/png"
	"log"
	"net/http"
	"rentWeb/kitex_gen/userapi"
	"rentWeb/utils"
	"time"
)


func GetSession(ginCtx *gin.Context){
	resp:=make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ginCtx.JSON(http.StatusOK,resp)
}

func GetImageID(ginCtx *gin.Context){
	//启动客户端
	c,err:=utils.GetUserServiceClient()
	if err != nil {
		log.Fatal(err)
	}
	uuid:=ginCtx.Param("uuid")
	//准备ImgCaptchaRequest
	req:=&userapi.ImgCaptchaRequest{
		Uuid: uuid,
	}

	resp,err:=c.GetImgCaptcha(context.Background(),req,callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	var img captcha.Image
	err=json.Unmarshal(resp.Img, &img)
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(ginCtx.Writer,img)
}

func GetSmsID(ginCtx *gin.Context) {
	c,err:=utils.GetUserServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	email := ginCtx.Param("email")
	imgCode := ginCtx.Query("text")
	uuid := ginCtx.Query("id")

	req := &userapi.SmsCaptchaRequest{
		Email:   email,
		ImgCode: imgCode,
		Uuid:    uuid,
	}

	resp, err := c.SendSmsCaptcha(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Errno== "0" {
		ginCtx.JSON(http.StatusOK, resp)
	}else {
		ginCtx.JSON(http.StatusBadRequest,resp)
	}
}

func PostRegister(ginCtx *gin.Context){
	/*from表单数据接收方式
	email:=ginCtx.PostForm("mobile")
	pwd:=ginCtx.PostForm("password")
	smsCode:=ginCtx.PostForm("sms_code")
	fmt.Println(email,pwd,smsCode)
	*/

	//ajax 传值用bind接受数据，在网页端体现为 request payload
	var regData struct{
		Email string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode string `json:"sms_code"`
	}
	ginCtx.Bind(&regData)

	c,err:=utils.GetUserServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	req:=&userapi.RegisterUserRequest{
		Email: regData.Email,
		Password: regData.PassWord,
		Smscode: regData.SmsCode,
	}
	resp,err:=c.RegisterUser(context.Background(),req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Errno=="0"{
		ginCtx.JSON(http.StatusOK,resp)
	}else {
		ginCtx.JSON(http.StatusBadRequest,resp)
	}
}