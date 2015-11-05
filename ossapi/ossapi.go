/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/log"
	"net/http"
)

//LogCat
var Logger *log.Logger

func init() {
	/*
		// for coverall 95%
		var err error
		Logger, err = log.NewFileLogger(".ossapilog", "ossapi")
		if err != nil {
			fmt.Errorf("Create Logger Error\n")
			return
		}
	*/
	Logger, _ = log.NewFileLogger(".ossapilog", "ossapi")
	Logger.SetMaxFileSize(1024 * 1024 * 100) //100MB
	Logger.SetLevel(log.LDEBUG)
}

var (
	// global Access Key ID
	accessKeyID string
	// global Access Key Secret
	accessKeySecret string
)

const (
	// DateFmt is format date
	DateFmt = "Mon, 02 Jan 2006 15:04:05 GMT"
)

// http client for http request
var httpClient http.Client

// Init ossapi with Access Key's ID and secret
// @param ID : Access Key's ID
// @param secret : Access Key's secret
func Init(ID string, secret string) error {
	if "" == ID || "" == secret {
		return ErrARG
	}
	accessKeyID = ID
	accessKeySecret = secret
	return nil
}
