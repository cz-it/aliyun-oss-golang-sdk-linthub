/**
* Author: CZ cz.theng@gmail.com
 */

package cors

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

// Query bucket's cors info
// @param bucketName : name of bucket
// @param location: location of buket
// @return rstinfo : CORS rules
// @rreturn ossapiError : nil on success
func Query(bucketName, location string) (rstInfo []CORSRuleInfo, ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName) + "/"
	req := &ossapi.Request{
		Host:     host,
		SubRes:   []string{"cors"},
		Path:     "?cors",
		Method:   "GET",
		Resource: resource}
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
	bodyLen, err := strconv.Atoi(rsp.HttpRsp.Header["Content-Length"][0])
	if err != nil {
		ossapi.Logger.Error("strconv.Atoi(rsp.HttpRsp.Header Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	rstBody := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(rstBody)
	info := new(CORSInfo)
	err = xml.Unmarshal(rstBody, info)
	if err != nil {
		ossapi.Logger.Error("xml.Unmarshal(body, rstInfo)Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	rstInfo = info.CORSRule
	return
}
