/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// LoggingInfo is Logging info
type LoggingInfo struct {
	TargetBucket string
	TargetPrefix string
}

// OpenLoggingInfo is  Open Loggign info
type OpenLoggingInfo struct {
	XMLName        xml.Name    `xml:"BucketLoggingStatus"`
	LoggingEnabled LoggingInfo `xml:"LoggingEnabled"`
}

// CloseLoggingInfo is  Close Logginginfo
type CloseLoggingInfo struct {
	XMLName xml.Name `xml:"BucketLoggingStatus"`
}

// OpenLogging  Open bucket's logging
// @param name: name of bucket
// @param locaton: location of bucket
// @param targetBucket: target bucket to store log
// @param targetPrefix: object's prefix name of logging file
// @return ossapiError: nil on success
func OpenLogging(name, location, targetBucket, targetPrefix string) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	info := LoggingInfo{
		TargetBucket: targetBucket,
		TargetPrefix: targetPrefix}
	openInfo := &OpenLoggingInfo{
		LoggingEnabled: info}
	body, err := xml.Marshal(openInfo)
	if err != nil {
		ossapi.Logger.Error("err := xml.Marshal(openInfo) Error %s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?logging",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"logging"},
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

// CloseLogging Close bucket's logging
// @param name: name of bucket
// @param locaton: location of bucket
// @return ossapiError: nil on success
func CloseLogging(name, location string) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	closeInfo := &CloseLoggingInfo{}
	body, err := xml.Marshal(closeInfo)
	if err != nil {
		ossapi.Logger.Error("err := xml.Marshal(closeInfo) Error %s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?logging",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"logging"},
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
