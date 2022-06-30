package main

import (
	"context"
	"homeService/kitex_gen/homeapi"
	"homeService/model"
)

// HomeServiceImpl implements the last service interface defined in the IDL.
type HomeServiceImpl struct{}

// GetAreas implements the HomeServiceImpl interface.
func (s *HomeServiceImpl) GetAreas(ctx context.Context, req *homeapi.AreaRequest) (resp *homeapi.AreaResponse, err error) {
	// TODO: Your code here...
	areaData,_:=model.GetAreaFromRedis()
	if areaData==nil||len(areaData)==0{
		areaData,_=model.GetAreaFromMysql()
	}
	resp=&homeapi.AreaResponse{
		areaData,
	}
	return resp,nil
}
