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

// FilterInfo is filter
type FilterInfo struct {
	Delimiter      string
	MaxUploads     int
	KeyMarker      string
	Prefix         string
	UploadIDMarker string
	Encoding       string
}

// UploadInfo is upload info
type UploadInfo struct {
	Key       string
	UploadId  string
	Initiated string
}

//MultipartUploadsResultInfo is result info
type MultipartUploadsResultInfo struct {
	XMLName            xml.Name `xml:"ListMultipartUploadsResult"`
	Bucket             string   `xml:"Bucket"`
	KeyMarker          string   `xml:"KeyMarker"`
	UploadIdMarker     string   `xml:"UploadIdMarker"`
	NextKeyMarker      string   `xml:"NextKeyMarker"`
	NextUploadIdMarker string   `xml:"NextUploadIdMarker"`
	Delimiter          string   `xml:"Delimiter"`
	Prefix             string   `xml:"Prefix"`
	MaxUploads         int      `xml:"MaxUploads"`
	IsTruncated        bool     `xml:"IsTruncated"`
	Upload             []UploadInfo
}

//QueryObjects query mutli object :w
func QueryObjects(bucketName, location string, filter *FilterInfo) (rstInfo *MultipartUploadsResultInfo, ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName)
	var args []string
	if filter != nil {
		if filter.Delimiter != "" {
			args = append(args, "delimiter="+filter.Delimiter)
		}
		if filter.Encoding != "" {
			args = append(args, "encoding-type="+filter.Encoding)
		}
		if filter.KeyMarker != "" {
			args = append(args, "key-marker="+filter.KeyMarker)
		}
		if filter.MaxUploads > 0 {
			args = append(args, "max-uploads="+strconv.FormatUint(uint64(filter.MaxUploads), 10))
		}
		if filter.Prefix != "" {
			args = append(args, "prefix="+filter.Prefix)
		}
		if filter.UploadIDMarker != "" {
			args = append(args, "upload-id-marker="+filter.UploadIDMarker)
		}
	}
	argsStr := ""
	if args != nil {
		argsStr = "&" + strings.Join(args, "&")
	}

	req := &ossapi.Request{
		Host:     host,
		Path:     "/?uploads" + argsStr,
		SubRes:   []string{"uploads"},
		Method:   "GET",
		Resource: resource + "/"}

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
	rstBody := make([]byte, bodyLen)
	rsp.HTTPRsp.Body.Read(rstBody)
	rstInfo = new(MultipartUploadsResultInfo)
	err = xml.Unmarshal(rstBody, rstInfo)
	if err != nil {
		ossapi.Logger.Error("xml.Unmarshal(rstBody, rstInfo)  Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
