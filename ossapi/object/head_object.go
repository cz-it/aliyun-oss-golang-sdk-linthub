/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

type ObjectBriefInfo struct {
	ObjectType  string
	Type        string
	LastModifed string
	ETag        string
	Length      uint64
}

type BriefConnInfo struct {
	ModifiedSince   string
	UnmodifiedSince string
	MatchEtag       string
	NotMatchEtag    string
}

func QueryMeta(objName, bucketName, location string, info *BriefConnInfo) (briefInfo *ObjectBriefInfo, ossapiError *ossapi.Error) {
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	header := make(map[string]string)
	if info != nil {
		if info.ModifiedSince != "" {
			header["If-Modified-Since"] = info.ModifiedSince
		}
		if info.UnmodifiedSince != "" {
			header["If-Unmodified-Since"] = info.UnmodifiedSince
		}
		if info.MatchEtag != "" {
			header["If-Match"] = info.MatchEtag
		}
		if info.NotMatchEtag != "" {
			header["If-None-Match"] = info.NotMatchEtag
		}
	}

	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + objName,
		Method:   "HEAD",
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
	length, err := strconv.Atoi(rsp.HttpRsp.Header["Content-Length"][0])
	if err != nil {
		length = 0
	}

	briefInfo = &ObjectBriefInfo{ObjectType: rsp.HttpRsp.Header["X-Oss-Object-Type"][0],
		Type:        rsp.HttpRsp.Header["Content-Type"][0],
		LastModifed: rsp.HttpRsp.Header["Last-Modified"][0],
		ETag:        rsp.HttpRsp.Header["Etag"][0],
		Length:      uint64(length),
	}
	return
}
