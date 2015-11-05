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

func TestQueryObject(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := Query("acl", "test-object-hz", bucket.L_Hangzhou, nil, nil); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("GetObjectACL Success!")
		fmt.Println(info)
	}

}
