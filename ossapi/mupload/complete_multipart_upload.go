// Package mupload

/**
* Author: CZ cz.theng@gmail.com
 */

package mupload

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

//PartInfo is part infomation
type PartInfo struct {
	PartNumber int
	ETag       string
}

//PartsInfo is a list partinfo
type PartsInfo struct {
	XMLName xml.Name `xml:"CompleteMultipartUpload"`
	Part    []PartInfo
}

//PartsCompleteInfo Parts complete Info
type PartsCompleteInfo struct {
	XMLName  xml.Name `xml:"CompleteMultipartUploadResult"`
	Location string   `xml:"Location"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	ETag     string   `xml:"ETag"`
}

// Complete is  Finish uploading
// @param objName: object's Name
// @param bucketName : bucket's name
// @param location: bucket's location
// @param uploadID: uploading context ID
// @param info : parts info
// @return rstInfo : return response
// @reurn ossapiError : nil on success
func Complete(objName, bucketName, location string, uploadID string, info *PartsInfo) (rstInfo *PartsCompleteInfo, ossapiError *ossapi.Error) {
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error(err.Error())
		ossapiError = ossapi.OSSAPIError
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + objName + "?uploadId=" + uploadID,
		Method:   "POST",
		Body:     body,
		CntType:  "application/xml",
		SubRes:   []string{"uploadId=" + uploadID},
		Resource: resource}

	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error(err.Error())
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
		ossapi.Logger.Error(err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	rstbody := make([]byte, bodyLen)
	rsp.HTTPRsp.Body.Read(rstbody)
	rstInfo = new(PartsCompleteInfo)
	err = xml.Unmarshal(body, rstInfo)
	if err != nil {
		ossapi.Logger.Error(err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
