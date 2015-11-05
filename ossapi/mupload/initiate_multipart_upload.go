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

// init info
type InitInfo struct {
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	Expires            string
	Encryption         string
}

// return response
type InitRstInfo struct {
	XMLName  xml.Name `xml:"InitiateMultipartUploadResult"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	UploadId string   `xml:"UploadId"`
}

// Init a uploading context
// @param objName: object's Name
// @param bucketName : bucket's name
// @param location: bucket's location
// @return rstInfo : uploading context info
// @reurn ossapiError : nil on success

func Init(objName, bucketName, location string, initInfo *InitInfo) (rstInfo *InitRstInfo, ossapiError *ossapi.Error) {
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	header := make(map[string]string)
	if initInfo != nil {
		header["Cache-Control"] = initInfo.CacheControl
		header["Content-Disposition"] = initInfo.ContentDisposition
		header["Content-Encoding"] = initInfo.ContentEncoding
		header["Expires"] = initInfo.Expires
	}
	req := &ossapi.Request{
		Host:      host,
		Path:      "/" + objName + "?uploads",
		Method:    "POST",
		Resource:  resource,
		SubRes:    []string{"uploads"},
		ExtHeader: header}
	req.AddXOSS("x-oss-server-side-encryption", initInfo.Encryption)

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
	rstInfo = new(InitRstInfo)
	err = xml.Unmarshal(body, rstInfo)
	if err != nil {
		ossapi.Logger.Error("xml.Unmarshal(body, rstInfo) Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
