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

type PartInfo struct {
	PartNumber int
	ETag       string
}

type PartsInfo struct {
	XMLName xml.Name `xml:"CompleteMultipartUpload"`
	Part    []PartInfo
}

type PartsCompleteInfo struct {
	XMLName  xml.Name `xml:"CompleteMultipartUploadResult"`
	Location string   `xml:"Location"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	ETag     string   `xml:"ETag"`
}

func Complete(objName, bucketName, location string, uploadId string, info *PartsInfo) (rstInfo *PartsCompleteInfo, ossapiError *ossapi.Error) {
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error("xml.Marshal(cfg) Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + objName + "?uploadId=" + uploadId,
		Method:   "POST",
		Body:     body,
		CntType:  "application/xml",
		SubRes:   []string{"uploadId=" + uploadId},
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
	rstbody := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(rstbody)
	rstInfo = new(PartsCompleteInfo)
	err = xml.Unmarshal(body, rstInfo)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
