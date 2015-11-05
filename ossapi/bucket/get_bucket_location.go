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

// LocationInfo
type LocationInfo struct {
	XMLName  xml.Name `xml:"LocationConstraint"`
	Location string   `xml:",chardata"`
}

// Query bucket's location
// @param name: name of bucket
// @return location : location name of bucket
// @return ossapiError : nil on success
func QueryLocation(name string) (location string, ossapiError *ossapi.Error) {
	host := name + ".oss.aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?location",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"location"}}
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
	locationInfo := new(LocationInfo)
	err = xml.Unmarshal(body, locationInfo)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	location = locationInfo.Location
	return
}
