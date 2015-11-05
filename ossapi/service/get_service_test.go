/**
* Author: CZ cz.theng@gmail.com
 */

package service

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestGetService(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if buckets, err := QueryBucketsDefault(); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println(buckets)
		t.Log("[SUCC]:GetService")
	}

	fmt.Println("+++++++++++++++Get Service With+++++++++++")
	if buckets, err := QueryBuckets("aa", "b&afds=safsd?asfsab", 10); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println(buckets)
		t.Log("[SUCC]:GetService")
	}

	fmt.Println("+++++++++++++++With Init Error+++++++++++")
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1sfasdfs") {
		t.Fail()
	}
	if buckets, err := QueryBucketsDefault(); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println(buckets)
		t.Log("[SUCC]:GetService")
	}
}
