/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// Set bucket's ACL info
// @param name : name of bucket
// @param permission : permisson to set
// @return ossapiError : nil on success
func SetACL(name, location, permission string) (error *ossapi.Error) {
	resource := path.Join("/", name)
	host := name + "." + location + ".aliyuncs.com"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?acl",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"acl"}}
	req.AddXOSS("x-oss-acl", permission)

	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
			error = ossapi.OSSAPIError
			return
		}
	}
	if rsp.Result != ossapi.ESUCC {
		error = err.(*ossapi.Error)
		return
	}
	b := make([]byte, 1024)
	rsp.HttpRsp.Body.Read(b)
	return
}
