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
// Redefine in put_bucket_website
type IndexInfo struct {
	Suffix string
}
type ErrorInfo struct {
	Key string
}
type WebsiteInfo struct {
	XMLName       xml.Name  `xml:"WebsiteConfiguration"`
	IndexDocument IndexInfo `xml:"IndexDocument"`
	ErrorDocument KeyInfo   `xml:"ErrorDocument"`
}
*/

//Query bucket's website info
// @param name: name of bucket
// @param location : location of bucket
// @return info : website info of bucket
// @return ossapiError : nil on success
func QueryWebsite(name, location string) (info *WebsiteInfo, ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?website",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"website"}}
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
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(body)
	info = new(WebsiteInfo)
	err = xml.Unmarshal(body, info)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
