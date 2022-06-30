// Code generated by Kitex v0.3.0. DO NOT EDIT.

package homeservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"homeService/kitex_gen/homeapi"
)

func serviceInfo() *kitex.ServiceInfo {
	return homeServiceServiceInfo
}

var homeServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "homeService"
	handlerType := (*homeapi.HomeService)(nil)
	methods := map[string]kitex.MethodInfo{
		"getAreas": kitex.NewMethodInfo(getAreasHandler, newHomeServiceGetAreasArgs, newHomeServiceGetAreasResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "homeapi",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.0",
		Extra:           extra,
	}
	return svcInfo
}

func getAreasHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*homeapi.HomeServiceGetAreasArgs)
	realResult := result.(*homeapi.HomeServiceGetAreasResult)
	success, err := handler.(homeapi.HomeService).GetAreas(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHomeServiceGetAreasArgs() interface{} {
	return homeapi.NewHomeServiceGetAreasArgs()
}

func newHomeServiceGetAreasResult() interface{} {
	return homeapi.NewHomeServiceGetAreasResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetAreas(ctx context.Context, req *homeapi.AreaRequest) (r *homeapi.AreaResponse, err error) {
	var _args homeapi.HomeServiceGetAreasArgs
	_args.Req = req
	var _result homeapi.HomeServiceGetAreasResult
	if err = p.c.Call(ctx, "getAreas", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}