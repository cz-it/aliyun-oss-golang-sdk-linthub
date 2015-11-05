/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"strings"
)

const (
	Major = "1"
	Minor = "0"
	Patch = "0"
)

func Version() string {
	return strings.Join([]string{Major, Minor, Patch}, ".")
}
