/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

//Key
type KeyInfo struct {
	Key string
}

//Delete Info
type DeleteObjInfo struct {
	XMLName xml.Name `xml:"Delete"`
	Quiet   bool     `xml:"Quiet"`
	Object  []KeyInfo
}

// Response info
type DeleteObjRstInfo struct {
	XMLName xml.Name  `xml:"DeleteResult"`
	Deleted []KeyInfo `xml:"Deleted"`
}

// Delte serveral object
// @param bucketName : name of bucket
// @param locaton : location of bucket
// @param info : list of objcets
// @return rstInfo : return deleted objects
// @return ossapiError: nil on success

func DeleteObjects(bucketName, location string, info *DeleteObjInfo) (rstInfo *DeleteObjRstInfo, ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName) + "/"
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error("err := xml.Marshal(Info) Error %s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?delete",
		Method:   "POST",
		Body:     body,
		SubRes:   []string{"delete"},
		CntType:  "application/xml",
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
	rspbody := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(rspbody)
	rstInfo = new(DeleteObjRstInfo)
	err = xml.Unmarshal(rspbody, rstInfo)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return

	}
	return
}
