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

func TestCopyObject(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}

	copyInfo := &CopyInfo{
		ObjectName: "test2.html",
		BucketName: "test-object-hz",
		Location:   bucket.L_Hangzhou,
		Source:     "/test-object-hz/test"}

	if info, err := Copy(copyInfo, nil); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("CopyObject Success")
		fmt.Println(info)
	}
	conInfo := &CopyConditionInfo{}
	if info, err := Copy(copyInfo, conInfo); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("CopyObject Success")
		fmt.Println(info)
	}
}
