/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

// const define
const (
	LifecycleStatsEnable  = "Enabled"
	LifecycleStatsDisable = "Disabled"
)

// Days
type ExpirationDaysInfo struct {
	Days uint
}

//Date
type ExpirationDateInfo struct {
	Date string
}

//Rule info
type RuleInfo struct {
	ID         string
	Prefix     string
	Status     string
	Expiration ExpirationDaysInfo
}

// Lifecycle XML info
type LifecycleConfiguration struct {
	XMLName xml.Name `xml:"LifecycleConfiguration"`
	Rule    []RuleInfo
}

// Set bucket's lifecycle
// @param  name : name of bucket
// @param location : location of bucket
// @param rules : rules to set
// @return ossapiError : nil on success
func SetLifecycle(name, location string, rules []RuleInfo) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	info := LifecycleConfiguration{Rule: rules}
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error("err := xml.Marshal(Info) Error %s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}

	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?lifecycle",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"lifecycle"},
		Body:     body,
		CntType:  "application/xml"}
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
