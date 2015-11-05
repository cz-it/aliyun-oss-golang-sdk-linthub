/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestSetACL(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if err := SetACL("test-put-bucket", L_Hangzhou, P_PublicRW); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("PutBucketACL Success!")
	}

}
