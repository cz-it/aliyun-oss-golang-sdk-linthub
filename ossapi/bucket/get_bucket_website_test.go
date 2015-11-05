/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestGetBucketWebsite(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := QueryWebsite("test-put-bucket4", L_Beijing); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("GetBucketWebsite Success")
		fmt.Println(info)
	}
}
