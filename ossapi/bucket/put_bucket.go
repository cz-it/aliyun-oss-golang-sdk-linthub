/**
* Author: CZ cz.theng@gmail.com
 */
// package bucket wrap opration for bucket
// such as create query delete modify and etc.

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// Location and Permission
const (
	L_Hangzhou      = "oss-cn-hangzhou"
	L_Shenzhen      = "oss-cn-shenzhen"
	L_Beijing       = "oss-cn-beijing"
	L_Qingdao       = "oss-cn-qingdao"
	L_Shanghai      = "oss-cn-shanghai"
	L_HongKong      = "oss-cn-hongkong"
	L_SiliconValley = "oss-us-west-1"
	L_Singapore     = "oss-ap-southeast-1"

	P_Private        = "private"
	P_PublicReadOnly = "public-read"
	P_PublicRW       = "public-read-write"
)

// Requestion's XML Content
type CreateBucketConfiguration struct {
	XMLName            xml.Name `xml:"CreateBucketConfiguration"`
	LocationConstraint string   `xml:"LocationConstraint"`
}

// Create Bucket with name/location and permission
// location is list above
// permission now can be three value
// @param name : name of bucket
// @param permission : permission of bucket . it is P_XXX
// @return ossapiError : nil on success
func Create(name, location, permission string) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	cfg := &CreateBucketConfiguration{LocationConstraint: location}
	body, err := xml.Marshal(cfg)
	if err != nil {
		ossapi.Logger.Error("xml.Marshal(cfg) Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
	}
	body = append([]byte(xml.Header), body...)
	resource := path.Join("/", name)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/",
		Method:   "PUT",
		Resource: resource + "/",
		Body:     body,
		CntType:  "application/xml"}
	req.AddXOSS("x-oss-acl", permission)

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
	return
}

// Create Bucket with default
func CreateDefault(name string) (ossapiError *ossapi.Error) {
	ossapiError = Create(name, L_Hangzhou, P_Private)
	return
}
