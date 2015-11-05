/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// Referer Listinfo
type RefererListInfo struct {
	Referer []string
}

// ReferCOnfigurationInfo
type RefererConfigurationInfo struct {
	XMLName           xml.Name        `xml:"RefererConfiguration"`
	AllowEmptyReferer bool            `xml:"AllowEmptyReferer"`
	RefererList       RefererListInfo `xml:"RefererList"`
}

//Set Referer of bucket
// @param name : name of bucket
// @param location: locaton of bucket
// @param enable : wheather allow white access
// @param url: urls list
// @return ossapiError: nil on success
func SetReferer(name, location string, enable bool, urls []string) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	refersInfo := RefererListInfo{Referer: urls}
	var info RefererConfigurationInfo
	if urls == nil {
		info = RefererConfigurationInfo{
			AllowEmptyReferer: enable}
	} else {
		info = RefererConfigurationInfo{
			AllowEmptyReferer: enable,
			RefererList:       refersInfo}
	}
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error("err := xml.Marshal(Info) Error %s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?referer",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"referer"},
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
	if rsp.Result != ossapi.ESUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}

	return
}
