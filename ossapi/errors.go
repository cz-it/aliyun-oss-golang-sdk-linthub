/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"errors"
)

var (
	//Inner Errors
	ESUCC    = errors.New("Success!")
	EFAIL    = errors.New("HTTP Request Failed(4xx/5xx)!")
	EUNKNOWN = errors.New("HTTP Request With Unknown Status (NOT 2xx/4xx/5xx)!")
	EARG     = errors.New("Invalied Argument!")
)

const (
	ENone                                = "None"
	EAPI                                 = "OSSAPISDK"
	EAccessDenied                        = "AccessDenied"
	EBucketAlreadyExists                 = "BucketAlreadyExists"
	EBucketNotEmpty                      = "BucketNotEmpty"
	EEntityTooLarge                      = "EntityTooLarge"
	EEntityTooSmall                      = "EntityTooSmall"
	EFileGroupTooLarge                   = "FileGroupTooLarge"
	EInvalidLinkName                     = "InvalidLinkName"
	ELinkPartNotExist                    = "LinkPartNotExist"
	EObjectLinkTooLarge                  = "ObjectLinkTooLarge"
	EFieldItemTooLong                    = "FieldItemTooLong"
	EFilePartInterity                    = "FilePartInterity"
	EFilePartNotExist                    = "FilePartNotExist"
	EFilePartStale                       = "FilePartStale"
	EIncorrectNumberOfFilesInPOSTRequest = "IncorrectNumberOfFilesInPOSTRequest"
	EInvalidArgument                     = "InvalidArgument"
	EInvalidAccessKeyId                  = "InvalidAccessKeyId"
	EInvalidBucketName                   = "InvalidBucketName"
	EInvalidDigest                       = "InvalidDigest"
	EInvalidEncryptionAlgorithmError     = "InvalidEncryptionAlgorithmError"
	EInvalidObjectName                   = "InvalidObjectName"
	EInvalidPart                         = "InvalidPart"
	EInvalidPartOrder                    = "InvalidPartOrder"
	EInvalidPolicyDocument               = "InvalidPolicyDocument"
	EInvalidTargetBucketForLogging       = "InvalidTargetBucketForLogging"
	EInternalError                       = "InternalError"
	EMalformedXML                        = "MalformedXML"
	EMalformedPOSTRequest                = "MalformedPOSTRequest"
	EMaxPOSTPreDataLengthExceededError   = "MaxPOSTPreDataLengthExceededError"
	EMethodNotAllowed                    = "MethodNotAllowed"
	EMissingArgument                     = "MissingArgument"
	EMissingContentLength                = "MissingContentLength"
	ENoSuchBucket                        = "NoSuchBucket"
	ENoSuchKey                           = "NoSuchKey"
	ENoSuchUpload                        = "NoSuchUpload"
	ENotImplemented                      = "NotImplemented"
	EPreconditionFailed                  = "PreconditionFailed"
	ERequestTimeTooSkewed                = "RequestTimeTooSkewed"
	ERequestTimeout                      = "RequestTimeout"
	ERequestIsNotMultiPartContent        = "RequestIsNotMultiPartContent"
	ESignatureDoesNotMatch               = "SignatureDoesNotMatch"
	ETooManyBuckets                      = "TooManyBuckets"
)

var (
	OSSAPIError = &Error{ErrNo: EAPI,
		ErrMsg:       "OSSAPI SDK's Inner Error,You Can Find More Details In Log Files",
		ErrDetailMsg: "OSSAPI SDK's Inner Error,You Can Find More Details In Log Files"}
	ArgError = &Error{ErrNo: EAPI,
		ErrMsg:       "Argument Error.Please Check Your Argment Such as nil",
		ErrDetailMsg: "Argument Error.Please Check Your Argment Such as nil"}
)

type Error struct {
	XMLName      xml.Name `xml:"Error"`
	ErrNo        string   `xml:"Code"`
	ErrMsg       string   `xml:"Message"`
	HttpStatus   int
	ErrDetailMsg string
}

func (e Error) Error() string {
	return e.ErrMsg
}
