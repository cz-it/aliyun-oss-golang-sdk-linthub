/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"testing"
)

func TestDoc(t *testing.T) {
	if 0 == doc() {
		t.Log("[PASS]:doc()")
	}
}
