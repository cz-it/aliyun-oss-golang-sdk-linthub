/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// Info is object's info
type Info struct {
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	Expires            string
	Encryption         string
	ACL                string
	Body               []byte
	Type               string
}

// Create an new object
// @param objName : name of object
// @param bucketName : name of bucket
// @param locaton : location of bucket
// @param objInfo : object's info
// @retun ossapiError : nil on success
func Create(objName, bucketName, location string, objInfo *Info) (ossapiError *ossapi.Error) {
	if objInfo == nil {
		ossapiError = ossapi.ArgError
		return
	}
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	header := make(map[string]string)
	if objInfo != nil {
		header["Cache-Control"] = objInfo.CacheControl
		header["Content-Disposition"] = objInfo.ContentDisposition
		header["Content-Encoding"] = objInfo.ContentEncoding
		header["Expires"] = objInfo.Expires
		header["x-oss-server-side-encryption"] = objInfo.Encryption
		header["x-oss-object-acl"] = objInfo.ACL
	}
	req := &ossapi.Request{
		Host:      host,
		Path:      "/" + objName,
		Method:    "PUT",
		Resource:  resource,
		Body:      objInfo.Body,
		CntType:   objInfo.Type,
		ExtHeader: header}
	req.AddXOSS("x-oss-object-acl", objInfo.ACL)
	req.AddXOSS("x-oss-server-side-encryption", objInfo.Encryption)

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
	return
}
