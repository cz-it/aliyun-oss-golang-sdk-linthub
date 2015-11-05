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
// redifine in put_bucket_logging.go
type LoggingInfo struct {
	TargetBucket string
	TargetPrefix string
}
*/

// LoggingStatus is logging status struct
type LoggingStatus struct {
	XMLName        xml.Name `xml:"BucketLoggingStatus"`
	LoggingEnabled LoggingInfo
}

// QueryLogging Query bucket's Logging info
//@param name: name of bucket
//@param location: location of bucket
//@return info : Logging info of bucket
//@return ossapiError : nil on success
func QueryLogging(name, location string) (info *LoggingInfo, ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?logging",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"logging"}}
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
	status := new(LoggingStatus)
	err = xml.Unmarshal(body, status)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	if status.LoggingEnabled.TargetBucket == "" {
		info = nil
	} else {
		info = &status.LoggingEnabled
	}
	return
}
