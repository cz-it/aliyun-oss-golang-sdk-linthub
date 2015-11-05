/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"testing"
)

func TestInit(t *testing.T) {
	if nil == Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Log("Init Success!")
	}
	if nil == Init("", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Log("Init Success!")
	}
}
