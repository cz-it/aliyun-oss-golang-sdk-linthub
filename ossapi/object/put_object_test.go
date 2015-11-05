/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
	"testing"
)

func TestCreateObject(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	objInfo := &Info{
		CacheControl:       "no-cache",
		ContentDisposition: "attachment;filename=oss_download.jpg",
		ContentEncoding:    "utf-8",
		Expires:            "Fri, 28 Feb 2012 05:38:42 GMT",
		Encryption:         "AES256",
		ACL:                bucket.PPrivate,
		Body:               []byte("<html><head></head><body>test</body></html>"),
		Type:               "text/html"}
	if err := Create("test", "test-object-hz", bucket.LHangzhou, objInfo); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("PutObject Success!")
	}

	if err := Create("test2", "test-object-hz", bucket.LHangzhou, objInfo); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("PutObject Success!")
	}
}
