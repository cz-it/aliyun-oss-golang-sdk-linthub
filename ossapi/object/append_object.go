/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

/*
//redefine on put_object
type ObjectInfo struct {
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	Expires            string
	Encryption         string
	ACL                string
	ObjName            string
	BucketName         string
	Location           string
	Body               []byte
	Type               string
}
*/

// Append Info
type AppendObjInfo struct {
	ObjectInfo
	Position uint64
}

// Resopnse Info
type AppendObjRspInfo struct {
	Possition uint64
	crc64     uint64
}

// Create a Appendable object
// @param objName : name of object
// @param bucketName : name of bucket
// @param locaton : location of bucket
// @param objInfo : object meta info
// @return rstInfo : possition and crc of data
// @retun ossapiError : nil on success
func Append(objName, bucketName, location string, objInfo *AppendObjInfo) (rstInfo *AppendObjRspInfo, ossapiError *ossapi.Error) {
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
		Path:      "/" + objName + "?append&position=" + strconv.FormatUint(objInfo.Position, 10),
		Method:    "POST",
		Resource:  resource,
		Body:      objInfo.Body,
		CntType:   objInfo.Type,
		SubRes:    []string{"append&position=" + strconv.FormatUint(objInfo.Position, 10)},
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
	if rsp.Result != ossapi.ESUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	pos, _ := strconv.Atoi(rsp.HttpRsp.Header["X-Oss-Next-Append-Position"][0])
	crc, _ := strconv.Atoi(rsp.HttpRsp.Header["X-Oss-Hash-Crc64ecma"][0])
	rstInfo = &AppendObjRspInfo{
		Possition: uint64(pos),
		crc64:     uint64(crc)}
	return
}
