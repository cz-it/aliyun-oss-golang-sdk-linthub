/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

//RspObjInfo is Response info
type RspObjInfo struct {
	CntType      string
	LastModified string
	ETag         string
	Ranges       string
	Type         string
	Length       int
	Data         []byte
}

// OverrideInfo is  Override option
type OverrideInfo struct {
	Type         string
	Language     string
	Expires      string
	CacheControl string
	Disposition  string
	Encoding     string
}

//ConditionInfo is  condition info
type ConditionInfo struct {
	Range        string
	LastModify   string
	LastUnModify string
	ETag         string
	ETagMatched  bool
}

// Query an object 's data or download
// @param objName : name of object
// @param bucketName : name of bucket
// @param locaton : location of bucket
// @param condInfo : condition to query
// @param overrideInfo : controller which to return
// @return data: object's data
// @retun ossapiError : nil on success
func Query(objName, bucketName, location string, condInfo *ConditionInfo, overrideInfo *OverrideInfo) (data []byte, ossapiError *ossapi.Error) {
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	header := make(map[string]string)
	if condInfo != nil {
		if condInfo.Range != "" {
			header["Range"] = condInfo.Range
		}
		if condInfo.LastModify != "" {
			header["If-Modified-Since"] = condInfo.LastModify
		}
		if condInfo.LastUnModify != "" {
			header["If-Unmodified-Since"] = condInfo.LastUnModify
		}
		if condInfo.ETagMatched && condInfo.ETag != "" {
			header["If-Match"] = condInfo.ETag
		}
		if !condInfo.ETagMatched && condInfo.ETag != "" {
			header["If-None-Match"] = condInfo.ETag
		}
	}

	overrideHeader := make(map[string]string)
	if overrideInfo != nil {
		if overrideInfo.CacheControl != "" {
			overrideHeader["response-cache-control"] = overrideInfo.CacheControl
		}
		if overrideInfo.Disposition != "" {
			overrideHeader["response-content-disposition"] = overrideInfo.Disposition
		}
		if overrideInfo.Encoding != "" {
			overrideHeader["response-content-encoding"] = overrideInfo.Encoding
		}
		if overrideInfo.Expires != "" {
			overrideHeader["response-expires"] = overrideInfo.Expires
		}
		if overrideInfo.Language != "" {
			overrideHeader["response-content-language"] = overrideInfo.Language
		}
		if overrideInfo.Type != "" {
			overrideHeader["response-content-type"] = overrideInfo.Type
		}
	}
	req := &ossapi.Request{
		Host:      host,
		Path:      "/" + objName,
		Method:    "GET",
		Resource:  resource,
		Override:  overrideHeader,
		ExtHeader: header}

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
	data = body
	return
}
