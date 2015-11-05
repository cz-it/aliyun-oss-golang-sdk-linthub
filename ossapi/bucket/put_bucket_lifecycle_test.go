/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestSetBucketLifecycle(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	rules := []RuleInfo{RuleInfo{Prefix: "nimei", Status: LifecycleStatsEnable, Expiration: ExpirationDaysInfo{2}}}
	if err := SetLifecycle("test-put-bucket3", L_Beijing, rules); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("SetBucketLiecycle Success")
	}
}
