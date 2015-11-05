/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// DeleteLifecycle delete Bucket's Lifecycle.
// All lifecycle rules will be delete
// @param name : bucket's Name
// @param location: bucket's Location
// @return : nil on Success else ossapi.Error
func DeleteLifecycle(name, location string) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?lifecycle",
		Method:   "DELETE",
		Resource: resource,
		SubRes:   []string{"lifecycle"}}
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
