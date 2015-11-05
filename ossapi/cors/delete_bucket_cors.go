/**
* Author: CZ cz.theng@gmail.com
 */

package cors

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// Delete bucket's cors attribute
// @param bucketName: name of bucket
// @param location : location of bucket
// @return ossapiError : nil on success
func Delete(bucketName, location string) (ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName) + "/"
	req := &ossapi.Request{
		Host:     host,
		SubRes:   []string{"cors"},
		Path:     "/?cors",
		Method:   "DELETE",
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
	return
}
