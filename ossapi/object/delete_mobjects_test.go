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

func TestDeleteObjects(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	info := &DeleteObjInfo{
		Quiet:  false,
		Object: []KeyInfo{KeyInfo{Key: "test"}, KeyInfo{Key: "test2"}},
	}
	if info, err := DeleteObjects("test-object-hz", bucket.L_Hangzhou, info); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("Delte Multiplie Objects Success!")
		fmt.Println(info)
	}

}
