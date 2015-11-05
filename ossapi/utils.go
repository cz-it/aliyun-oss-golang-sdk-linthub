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

// Base64AndHmacSha1 caculate base64 and sha
func Base64AndHmacSha1(key, data []byte) (encStr string, err error) {
	if nil == data {
		err = ErrARG
		return
	}
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	expectedMAC := mac.Sum(nil)
	encStr = strings.TrimSpace(base64.StdEncoding.EncodeToString(expectedMAC))
	return
}

// Base64AndMd5 caclute Base64 and md5
func Base64AndMd5(data []byte) (encStr string, err error) {
	if nil == data {
		err = ErrARG
		return
	}
	md5Digest := md5.Sum(data)
	base64Str := base64.StdEncoding.EncodeToString(md5Digest[:])
	encStr = strings.TrimSpace(base64Str)
	return
}
