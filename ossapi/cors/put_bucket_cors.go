/**
* Author: CZ cz.theng@gmail.com
 */

package cors

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

//RuleInfo is CORS rules
type RuleInfo struct {
	AllowedOrigin []string
	AllowedMethod []string
	AllowedHeader []string
	ExposeHeader  []string
	MaxAgeSeconds uint64
}

// Info is  XML wraper
type Info struct {
	XMLName  xml.Name `xml:"CORSConfiguration"`
	CORSRule []RuleInfo
}

// Create Create a CORS rule
// @param bucketName : name of bucket
// @param location : bucket's loction
// @param corsInfo : cors rules
// @return ossapiError : nil on success
func Create(bucketName, location string, corsInfo []RuleInfo) (ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	info := &Info{CORSRule: corsInfo}
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error("xml.Marshal(cfg) Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	resource := path.Join("/", bucketName)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?cors",
		Method:   "PUT",
		SubRes:   []string{"cors"},
		Resource: resource + "/",
		Body:     body,
		CntType:  "application/xml"}

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
	return
}
