/**
* Author: CZ cz.theng@gmail.com
 */

package cors

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

type OptionReqInfo struct {
	Origin  string
	Method  string
	Headers string
}

type OptionRspInfo struct {
	AllowOrigin   string
	AllowMethods  string
	AllowHeaders  string
	ExposeHeaders string
	MaxAge        uint64
}

func Option(objName, bucketName, location string, optionInfo *OptionReqInfo) (rstInfo *OptionRspInfo, ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName, objName)
	headers := make(map[string]string)
	if optionInfo != nil {
		headers["Origin"] = optionInfo.Origin
		headers["Access-Control-Request-Method"] = optionInfo.Method
		headers["Access-Control-Request-Headers"] = optionInfo.Headers
	}
	req := &ossapi.Request{
		ExtHeader: headers,
		Host:      host,
		Path:      "/" + objName,
		Method:    "OPTIONS",
		Resource:  resource}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = ossapi.OSSAPIError
			return
		}
	}
	if rsp.Result != ossapi.ESUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	rstInfo = new(OptionRspInfo)
	rstInfo.AllowOrigin = rsp.HttpRsp.Header["Access-Control-Allow-Origin"][0]
	rstInfo.AllowMethods = rsp.HttpRsp.Header["Access-Control-Allow-Methods"][0]
	rstInfo.ExposeHeaders = rsp.HttpRsp.Header["Access-Control-Expose-Headers"][0]
	rstInfo.AllowHeaders = rsp.HttpRsp.Header["Access-Control-Allow-Headers"][0]
	age, _ := strconv.Atoi(rsp.HttpRsp.Header["Access-Control-Max-Age"][0])
	rstInfo.MaxAge = uint64(age)
	return
}
