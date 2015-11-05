/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	ArgError.Error()
}

func TestBase64AndHmacSha1(t *testing.T) {
	Base64AndHmacSha1([]byte("b"), []byte("bb"))
	Base64AndHmacSha1(nil, nil)
}

func TestBase64AndMd5(t *testing.T) {
	Base64AndMd5([]byte("abc"))
	Base64AndMd5(nil)
}

func TestDo(t *testing.T) {
	headers := map[string]string{"Content-XX": "adf"}
	req := &Request{
		Host:      "oss.aliyuncs.com",
		Path:      "/",
		ExtHeader: headers,
		SubRes:    []string{"aa"},
		Body:      []byte("aaa"),
		CntType:   "text/html",
		Method:    "GET",
		Resource:  "/"}
	req.AddXOSS("oss-xx", "abc")
	rsp, err := req.Send()

	fmt.Println("====Method Error=======")
	req.Method = "error"
	rsp, err = req.Send()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====URL Error=======")
	req.Host = "nimeidenimei"
	rsp, err = req.Send()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====New Request Error=======")
	req.Host = "//?a=b/tcp://abc:udp:"
	req.Path = ""
	rsp, err = req.Send()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====404 Error=======")
	req.Host = "www.baidu.com?cz=cz"
	rsp, err = req.Send()
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)

	fmt.Println(Version())
}
