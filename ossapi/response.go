/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"net/http"
)

//Response to Request
type Response struct {
	Result  error
	HTTPRsp *http.Response
}

// Tag Info for a Response
type Tag struct {
	RequestID string
	ETag      string
	Date      string
	Server    string
}
