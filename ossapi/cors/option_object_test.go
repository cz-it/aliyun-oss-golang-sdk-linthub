/**
* Author: CZ cz.theng@gmail.com
 */
package cors

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
	"testing"
)

func TestOptionObject(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	optionInfo := &OptionReqInfo{Origin: "www.qq.com", Method: "GET", Headers: "authorization"}
	if info, err := Option("app.py", "test-cors", bucket.LHangzhou, optionInfo); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("OptionObject Success")
		fmt.Println(info)
	}
}
