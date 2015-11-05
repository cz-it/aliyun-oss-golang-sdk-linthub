/**
* Author: CZ cz.theng@gmail.com
 */

package cors

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

// OptionReqInfo is Reqinfo
type OptionReqInfo struct {
	Origin  string
	Method  string
	Headers string
}

// OptionRspInfo is  Resoponse info
type OptionRspInfo struct {
	AllowOrigin   string
	AllowMethods  string
	AllowHeaders  string
	ExposeHeaders string
	MaxAge        uint64
}

// Option  Query CORS permission of bucket
// @param objName : object to access
// @param bucketName: bucket to access
// @param location: bucket's location
// @param optionInfo : CORS requet
// @return rstInfo: CORS permisson
// @return ossapiError : nil on success
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
	if rsp.Result != ossapi.ErrSUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	rstInfo = new(OptionRspInfo)
	rstInfo.AllowOrigin = rsp.HTTPRsp.Header["Access-Control-Allow-Origin"][0]
	rstInfo.AllowMethods = rsp.HTTPRsp.Header["Access-Control-Allow-Methods"][0]
	rstInfo.ExposeHeaders = rsp.HTTPRsp.Header["Access-Control-Expose-Headers"][0]
	rstInfo.AllowHeaders = rsp.HTTPRsp.Header["Access-Control-Allow-Headers"][0]
	age, _ := strconv.Atoi(rsp.HTTPRsp.Header["Access-Control-Max-Age"][0])
	rstInfo.MaxAge = uint64(age)
	return
}
