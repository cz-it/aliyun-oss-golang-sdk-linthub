// Package cors is for cors
/**
* Author: CZ cz.theng@gmail.com
 */
/**
CORS allow remote domain to access resources. Before this ,you should create a CORS rule for the bucket.

cors.Create

Create(bucketName, location string, corsInfo []CORSRuleInfo) (ossapiError *ossapi.Error)
Create CORS attribute for bucket.OSCR info is descriped by CORSRuleInfo.

type CORSRuleInfo struct {
    AllowedOrigin []string
    AllowedMethod []string
    AllowedHeader []string
    ExposeHeader  []string
    MaxAgeSeconds uint64
}
AllowedOrigin: origin vister source .
AllowedMethod: Valied Http Method
AllowedHeader: Valied Http Header
ExposeHeader : Valied header from client
MaxAgeSeconds： Cache time
Info is a Rule list. With AllowedOrigin AllowedMethod AllowedHeader ExposeHeader MaxAgeSeconds

cors.Query

Query(bucketName, location string) (rstInfo []CORSRuleInfo, ossapiError *ossapi.Error)
Query CORS information of bucket. CORS is store in CORSRuleInfo which descript above.

Info is a Rule list. With AllowedOrigin AllowedMethod AllowedHeader ExposeHeader MaxAgeSeconds

cors.Delete

Delete(bucketName, location string) (ossapiError *ossapi.Error)
Delete CORS infomation of bucket.After this, other domain can't vister this bucket.

Info is a Rule list. With AllowedOrigin AllowedMethod AllowedHeader ExposeHeader MaxAgeSeconds

cors.Option

Option(objName, bucketName, location string, optionInfo *OptionReqInfo) (rstInfo *OptionRspInfo, ossapiError *ossapi.Error)
Query Bucket's CORS information to decide weather visitable. vister should set OptionReqInfo:

type OptionReqInfo struct {
    Origin  string
    Method  string
    Headers string
}
CORS information is stored in OptionRspInfo:

type OptionRspInfo struct {
    AllowOrigin   string
    AllowMethods  string
    AllowHeaders  string
    ExposeHeaders string
    MaxAge        uint64
}
Info is a Rule list. With AllowedOrigin AllowedMethod AllowedHeader ExposeHeader MaxAgeSeconds
 **/
package cors

import ()
