/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

// OwnerInfo
type OwnerInfo struct {
	ID          string
	DisplayName string
}

//AccessControlListInfo
type AccessControlListInfo struct {
	Grant string
}

// ACLInfo
type ACLInfo struct {
	XMLName           xml.Name `xml:"AccessControlPolicy"`
	Owner             OwnerInfo
	AccessControlList AccessControlListInfo
}

//Query bucket's ACL
// @param name: name of bucket
// @param location: location of bucket
// @return info: ACL info of bucket
// @return ossapiError: nil on success
func QueryACL(name, location string) (info *ACLInfo, ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?acl",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"acl"}}
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
	info = new(ACLInfo)
	xml.Unmarshal(body, info)
	return
}
