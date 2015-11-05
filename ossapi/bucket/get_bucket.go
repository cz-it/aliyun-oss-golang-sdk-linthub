/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/service"
	"path"
	"strconv"
	"strings"
)

/*
// redefined in get_service.go
type Owner struct {
	ID          string
	DisplayName string
}
*/
type ContentInfo struct {
	Key          string
	LastModified string
	ETag         string
	Type         string
	Size         string
	StorageClass string
	Owner        service.Owner
}

// CommonInfo
type CommonInfo struct {
	Prefix string
}

//BucketsInfo
type BucktsInfo struct {
	XMLName        xml.Name `xml:"ListBucketResult"`
	Name           string   `xml:"Name"`
	Prefix         string   `xml:"Prefix"`
	Marker         string   `xml:"Marker"`
	MaxKeys        int      `xml:"MaxKeys"`
	EncodingType   string   `xml:"encoding-type"`
	IsTruncated    bool     `xml:"IsTruncated"`
	Contents       []ContentInfo
	CommonPrefixes CommonInfo `xml:"CommonPrefixes"`
}

// Query all objects of a bucket
// @param name : name of bucket
// @param location: location of bucket
// @param prefix: select valied prefix
// @param marker: marker after this will be return
// @param delimiter: valied delimiter, common prefix
// @param encoding: encoding of content
// @param maxKeys : at most maxKeys items will return
// @return info : objects' info
// @return ossapiError : nil on success
func QueryObjects(name, location string, prefix, marker, delimiter, encodingType string, maxKeys int) (info *BucktsInfo, ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	urlPath := "/"
	var args []string
	if prefix != "" {
		args = append(args, "prefix="+prefix)
	}
	if marker != "" {
		args = append(args, "marker="+marker)
	}
	if delimiter != "" {
		args = append(args, "delimiter="+delimiter)
	}
	if encodingType != "" {
		args = append(args, "encoding-type"+encodingType)
	}
	if maxKeys > 0 {
		args = append(args, "max-keys="+strconv.FormatUint(uint64(maxKeys), 10))
	}
	if args != nil {
		urlPath += "?" + strings.Join(args, "&")
	}
	req := &ossapi.Request{
		Host:     host,
		Path:     urlPath,
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
	body := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(body)
	info = new(BucktsInfo)
	err = xml.Unmarshal(body, info)
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	return
}
