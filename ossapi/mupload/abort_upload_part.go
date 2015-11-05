/**
* Author: CZ cz.theng@gmail.com
 */

package mupload

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// Abort uploading
// @param objName: object's Name
// @param bucketName : bucket's name
// @param location: bucket's location
// @param uploadID: uploading context ID
// @reurn ossapiError : nil on success
func Abort(objName, bucketName, location, uploadID string) (ossapiError *ossapi.Error) {
	host := bucketName + "." + location + ".aliyuncs.com"
	resource := path.Join("/", bucketName, objName)
	urlPath := "/" + objName + "?uploadId=" + uploadID
	req := &ossapi.Request{
		Host:     host,
		Path:     urlPath,
		SubRes:   []string{"uploadId=" + uploadID},
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
		fmt.Println(ossapiError.ErrDetailMsg)
		return
	}
	return
}
