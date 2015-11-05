/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

//Index info
type IndexInfo struct {
	Suffix string
}

// Key Info
type KeyInfo struct {
	Key string
}

// Website info
type WebsiteInfo struct {
	XMLName       xml.Name  `xml:"WebsiteConfiguration"`
	IndexDocument IndexInfo `xml:"IndexDocument"`
	ErrorDocument KeyInfo   `xml:"ErrorDocument"`
}

// Set bucket's website
// @param name: name of bucket
// @param location: location of bucket
// @param indexPage : index page
// @param errorPage : 404 error page
// @return ossapiError : nil on success
func SetWebsite(name, location, indexPage, errorPage string) (ossapiError *ossapi.Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	indexInfo := IndexInfo{Suffix: indexPage}
	var keyInfo KeyInfo
	keyInfo = KeyInfo{Key: errorPage}
	var info WebsiteInfo
	if "" == errorPage {
		info = WebsiteInfo{IndexDocument: indexInfo}
	} else {
		info = WebsiteInfo{IndexDocument: indexInfo, ErrorDocument: keyInfo}
	}
	body, err := xml.Marshal(info)
	if err != nil {
		ossapi.Logger.Error("err := xml.Marshal(Info) Error %s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &ossapi.Request{
		Host:     host,
		Path:     "/?website",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"website"},
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
