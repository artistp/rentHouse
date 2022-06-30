package controller

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rentWeb/entity"
	"rentWeb/kitex_gen/homeapi"
	"rentWeb/utils"
)


func GetArea(ginCtx *gin.Context){
	c,err:=utils.GetHomeServiceClient()
	if err != nil {
		log.Println()
		return
	}

	req:=&homeapi.AreaRequest{}

	remoteResp,_:=c.GetAreas(context.Background(),req)

	var areas []entity.Area
	json.Unmarshal(remoteResp.Areas,&areas)

	resp:=make(map[string]interface{})

	resp["errno"]="0"
	resp["errmsg"]=utils.RecodeText(utils.RECODE_OK)
	resp["data"]=areas

	ginCtx.JSON(http.StatusOK,resp)

}


