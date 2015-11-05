/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestSetBucketReferer(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if err := SetReferer("test-put-bucket3", L_Beijing, false, []string{"http://www.baidu.com", "http://www.qq.com"}); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("SetBucketReferer Success")
	}
	if err := SetReferer("test-put-bucket4", L_Beijing, false, nil); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("SetBucketReferer Success")
	}
}
