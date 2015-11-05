/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

func Base64AndHmacSha1(key, data []byte) (encStr string, err error) {
	if nil == data {
		err = EARG
		return
	}
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	expectedMAC := mac.Sum(nil)
	encStr = strings.TrimSpace(base64.StdEncoding.EncodeToString(expectedMAC))
	return
}

func Base64AndMd5(data []byte) (encStr string, err error) {
	if nil == data {
		err = EARG
		return
	}
	md5Digest := md5.Sum(data)
	base64Str := base64.StdEncoding.EncodeToString(md5Digest[:])
	encStr = strings.TrimSpace(base64Str)
	return
}
