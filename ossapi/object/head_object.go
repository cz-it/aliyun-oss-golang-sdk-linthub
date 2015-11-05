/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

// BriefInfo is  Brief Info
type BriefInfo struct {
	ObjectType  string
	Type        string
	LastModifed string
	ETag        string
	Length      uint64
}

// BriefConnInfo is Conndtition info
type BriefConnInfo struct {
	ModifiedSince   string
	UnmodifiedSince string
	MatchEtag       string
	NotMatchEtag    string
}

// QueryMeta  Query object's meta info
// @param objName : name of object
// @param bucketName : name of bucket
// @param locaton : location of bucket
// @param info: conndtion to controll return
// @return briefInfo : breif info of object
// @retun ossapiError : nil on success
func QueryMeta(objName, bucketName, location string, info *BriefConnInfo) (briefInfo *BriefInfo, ossapiError *ossapi.Error) {
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
	if rsp.Result != ossapi.ErrSUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	length, err := strconv.Atoi(rsp.HTTPRsp.Header["Content-Length"][0])
	if err != nil {
		length = 0
	}

	briefInfo = &BriefInfo{ObjectType: rsp.HTTPRsp.Header["X-Oss-Object-Type"][0],
		Type:        rsp.HTTPRsp.Header["Content-Type"][0],
		LastModifed: rsp.HTTPRsp.Header["Last-Modified"][0],
		ETag:        rsp.HTTPRsp.Header["Etag"][0],
		Length:      uint64(length),
	}
	return
}
