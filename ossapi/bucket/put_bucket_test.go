/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestCreate(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if err := CreateDefault("test-put-bucket2"); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println("Create Default Success!")
	}
	if err := Create("test-put-bucket3", L_Beijing, P_Private); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println("Create Success")
	}
}
