/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"strings"
)

const (
	major = "1"
	minor = "0"
	patch = "0"
)

// Version show OSSAPI's verison
func Version() string {
	return strings.Join([]string{major, minor, patch}, ".")
}
