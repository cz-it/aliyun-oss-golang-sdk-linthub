/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestGetBucketReferer(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := QueryReferer("test-put-bucket3", LBeijing); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("GetBucketReferer Success")
		fmt.Println(info)
	}
}
