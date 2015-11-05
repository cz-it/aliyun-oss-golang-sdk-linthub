/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

/*
//redefine in put_bucket_referer
type RefererListInfo struct {
	Referer []string
}

type RefererConfigurationInfo struct {
	XMLName           xml.Name        `xml:"RefererConfiguration"`
	AllowEmptyReferer bool            `xml:"AllowEmptyReferer"`
	RefererList       RefererListInfo `xml:"RefererList"`
}
*/

// QueryReferer Query Referer info of buckets
// @param name : name of bucket
// @param location: location of bucket
// @return info : referer info of bucket
// @return ossapiError : nil on success
func QueryReferer(name, location string) (info *RefererConfigurationInfo, ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?referer",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"referer"}}
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
	bodyLen, err := strconv.Atoi(rsp.HTTPRsp.Header["Content-Length"][0])
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.HTTPRsp.Body.Read(body)
	info = new(RefererConfigurationInfo)
	err = xml.Unmarshal(body, info)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
