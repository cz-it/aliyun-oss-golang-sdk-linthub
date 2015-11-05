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

func TestCreateBucketCORS(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	corsInfo := []CORSRuleInfo{CORSRuleInfo{
		AllowedOrigin: []string{"www.qq.com", "www.baidu.com", "*", "/"},
		AllowedMethod: []string{"GET"},
		AllowedHeader: []string{"Access-Control-Requet-Headers", "Authorization", "Access-Control-Request-Method"},
		ExposeHeader:  []string{"x-oss-test", "x-oss-test2"},
		MaxAgeSeconds: 100,
	}}
	if err := Create("test-cors", bucket.L_Hangzhou, corsInfo); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println("Create Bucket CORS SUCCESS")
	}
}
