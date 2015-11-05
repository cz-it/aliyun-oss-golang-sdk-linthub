/**
* Author: CZ cz.theng@gmail.com
 */

package mupload

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
	"strings"
)

// Parts info
type PartsFilterInfo struct {
	MaxParts         int
	PartNumberMarker int
	Encoding         string
}

// Parts list
type PartListInfo struct {
	PartNumber   int
	LastModified string
	ETag         string
	Size         uint64
}

// Return Response
type PartsResultInfo struct {
	XMLName              xml.Name `xml:"ListPartsResult"`
	Bucket               string   `xml:"Bucket"`
	Key                  string   `xml:"Key"`
	UploadId             string   `xml:"UploadId"`
	NextPartNumberMarker string   `xml:"NextPartNumberMarker"`
	MaxParts             int      `xml:"MaxParts"`
	IsTruncated          bool     `xml:"IsTruncated"`
	Part                 []PartListInfo
}

// Qeury uploaded parts info
// @param objName: object's Name
// @param bucketName : bucket's name
// @param location: bucket's location
// @param uploadID: uploading context ID
// @param filter: filter to query
// @return rstInfo: return parts
// @reurn ossapiError : nil on success

func QueryParts(objName, bucketName, location string, uploadID string, filter *PartsFilterInfo) (rstInfo *PartsResultInfo, ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName, objName)
	var args []string
	if uploadID != "" {
		args = append(args, "uploadId="+uploadID)
	}
	if filter != nil {
		if filter.Encoding != "" {
			args = append(args, "=encoding-type"+filter.Encoding)
		}
		if filter.MaxParts > 0 {
			args = append(args, "max-parts="+strconv.FormatUint(uint64(filter.MaxParts), 10))
		}
		if filter.PartNumberMarker > 0 {
			args = append(args, "part-number-marker="+strconv.FormatUint(uint64(filter.PartNumberMarker), 10))
		}
	}
	argsStr := ""
	if args != nil {
		argsStr = strings.Join(args, "&")
	}

	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + objName + "?" + argsStr,
		SubRes:   []string{"uploadId=" + uploadID},
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
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	rstBody := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(rstBody)
	rstInfo = new(PartsResultInfo)
	err = xml.Unmarshal(rstBody, rstInfo)
	if err != nil {
		ossapi.Logger.Error("xml.Unmarshal(rstBody, rstInfo)  Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
