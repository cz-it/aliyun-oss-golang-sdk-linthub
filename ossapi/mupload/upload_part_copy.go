/**
* Author: CZ cz.theng@gmail.com
 */

package mupload

import (
	"encoding/xml"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/object"
	"path"
	"strconv"
)

type UploadPartCopyInfo struct {
	ObjectName string
	BucketName string
	Location   string
	PartNumber int
	UploadID   string
	//	Data          []byte
	//	CntType       string
	SrcObjectName string
	SrcBucketName string
	SrcRangeBegin int
	SrcRangeEnd   int
}

type UploadPartCopyRstInfo struct {
	XMLName      xml.Name `xml:"CopyObjectResult"`
	LastModified string   `xml:"LastModified"`
	ETag         string   `xml:"ETag"`
}

func Copy(partInfo *UploadPartCopyInfo, copyConnInfo *object.CopyConditionInfo) (rstInfo *UploadPartCopyRstInfo, ossapiError *ossapi.Error) {
	if partInfo == nil {
		ossapiError = ossapi.ArgError
		return
	}
	resource := path.Join("/", partInfo.BucketName, partInfo.ObjectName)
	host := partInfo.BucketName + "." + partInfo.Location + ".aliyuncs.com"
	header := make(map[string]string)
	header["Content-Length"] = strconv.FormatUint(uint64(partInfo.SrcRangeEnd-partInfo.SrcRangeBegin), 10)
	req := &ossapi.Request{
		Host:      host,
		ExtHeader: header,
		Path:      "/" + partInfo.ObjectName + "?partNumber=" + strconv.FormatUint(uint64(partInfo.PartNumber), 10) + "uploadId=" + partInfo.UploadID,
		Method:    "PUT",
		//		Body:     partInfo.Data,
		//		CntType:  partInfo.CntType,
		SubRes:   []string{"partNumber=" + strconv.FormatUint(uint64(partInfo.PartNumber), 10) + "uploadId=" + partInfo.UploadID},
		Resource: resource}
	if copyConnInfo != nil {
		req.AddXOSS("x-oss-copy-source-if-match", copyConnInfo.ETAG)
		req.AddXOSS("x-oss-copy-source-if-none-match", copyConnInfo.Date)
		req.AddXOSS("x-oss-copy-source-if-unmodified-since", copyConnInfo.LastUnModify)
		req.AddXOSS("x-oss-copy-source-if-modified-since", copyConnInfo.LastModify)
	}
	if partInfo.SrcObjectName != "" && partInfo.SrcBucketName != "" {
		req.AddXOSS("x-oss-copy-source", path.Join("/", partInfo.SrcBucketName, partInfo.SrcObjectName))
	}

	if partInfo.SrcRangeBegin > 0 && partInfo.SrcRangeEnd > 0 {
		cntRange := "bytes=" + strconv.FormatUint(uint64(partInfo.SrcRangeBegin), 10) + "-" + strconv.FormatUint(uint64(partInfo.SrcRangeEnd), 10)
		req.AddXOSS("x-oss-copy-source-range", cntRange)
	}

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
	fmt.Println("rsp:", rsp.HttpRsp)
	bodyLen, err := strconv.Atoi(rsp.HttpRsp.Header["Content-Length"][0])
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(body)
	fmt.Println("rsp:body", string(body))
	rstInfo = new(UploadPartCopyRstInfo)
	xml.Unmarshal(body, rstInfo)
	return
}
