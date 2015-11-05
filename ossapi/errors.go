//Package ossapi is oss sdk of golang
/**
* Author: CZ cz.theng@gmail.com
 */
package ossapi

import (
	"encoding/xml"
	"errors"
)

var (
	//ErrSUCC is none error
	ErrSUCC = errors.New("Success!")
	//ErrFAIL is all fail error
	ErrFAIL = errors.New("HTTP Request Failed(4xx/5xx)")
	// ErrUNKNOWN is other errors
	ErrUNKNOWN = errors.New("HTTP Request With Unknown Status (NOT 2xx/4xx/5xx)")
	//ErrARG is args error such as nil
	ErrARG = errors.New("Invalied Argument!")
)

const (
	// ErrNone is none error
	ErrNone = "None"
	//EAPI is SDK API's error
	EAPI = "OSSAPISDK"
	// EAccessDenied is oss's EAccessDenied
	EAccessDenied = "AccessDenied"
	// EBucketAlreadyExists is oss's EBucketAlreadyExists
	EBucketAlreadyExists = "BucketAlreadyExists"
	// EBucketNotEmpty is oss's EBucketNotEmpty
	EBucketNotEmpty = "BucketNotEmpty"
	// EEntityTooLarge is oss's EEntityTooLarge
	EEntityTooLarge = "EntityTooLarge"
	// EEntityTooSmall is oss's EEntityTooSmall
	EEntityTooSmall = "EntityTooSmall"
	// EFileGroupTooLarge is oss's EFileGroupTooLarge
	EFileGroupTooLarge = "FileGroupTooLarge"
	//EInvalidLinkName is oss's EInvalidLinkName
	EInvalidLinkName = "InvalidLinkName"
	// ELinkPartNotExist is oss's ELinkPartNotExist
	ELinkPartNotExist = "LinkPartNotExist"
	// EObjectLinkTooLarge is oss's EObjectLinkTooLarge
	EObjectLinkTooLarge = "ObjectLinkTooLarge"
	// EFieldItemTooLong is oss's EFieldItemTooLong
	EFieldItemTooLong = "FieldItemTooLong"
	//EFilePartInterity is oss's EFilePartInterity
	EFilePartInterity = "FilePartInterity"
	//EFilePartNotExist is oss's EFilePartNotExist
	EFilePartNotExist = "FilePartNotExist"
	//EFilePartStale is oss' EFilePartStale
	EFilePartStale = "FilePartStale"
	// EIncorrectNumberOfFilesInPOSTRequest is oss's EIncorrectNumberOfFilesInPOSTRequest
	EIncorrectNumberOfFilesInPOSTRequest = "IncorrectNumberOfFilesInPOSTRequest"
	// EInvalidArgument is oss's EInvalidArgument
	EInvalidArgument = "InvalidArgument"
	// EInvalidAccessKeyID is oss's EInvalidAccessKeyId
	EInvalidAccessKeyID = "InvalidAccessKeyId"
	// EInvalidBucketName is oss's EInvalidBucketName
	EInvalidBucketName = "InvalidBucketName"
	//EInvalidDigest is oss's EInvalidDigest
	EInvalidDigest = "InvalidDigest"
	//EInvalidEncryptionAlgorithmError is oss's EInvalidEncryptionAlgorithmError
	EInvalidEncryptionAlgorithmError = "InvalidEncryptionAlgorithmError"
	//EInvalidObjectName is oss' EInvalidObjectName
	EInvalidObjectName = "InvalidObjectName"
	// EInvalidPart is oss's EInvalidPart
	EInvalidPart = "InvalidPart"
	// EInvalidPartOrder is oss's EInvalidPartOrder
	EInvalidPartOrder = "InvalidPartOrder"
	// EInvalidPolicyDocument is oss's EInvalidPolicyDocument
	EInvalidPolicyDocument = "InvalidPolicyDocument"
	// EInvalidTargetBucketForLogging is oss's EInvalidTargetBucketForLogging
	EInvalidTargetBucketForLogging = "InvalidTargetBucketForLogging"
	// EInternalError is oss's EInternalError
	EInternalError = "InternalError"
	// EMalformedXML is oss's EMalformedXML
	EMalformedXML = "MalformedXML"
	// EMalformedPOSTRequest is oss's EMalformedPOSTRequest
	EMalformedPOSTRequest = "MalformedPOSTRequest"
	// EMaxPOSTPreDataLengthExceededError is oss's EMaxPOSTPreDataLengthExceededError
	EMaxPOSTPreDataLengthExceededError = "MaxPOSTPreDataLengthExceededError"
	//EMethodNotAllowed is oss's EMethodNotAllowed
	EMethodNotAllowed = "MethodNotAllowed"
	//EMissingArgument is oss's EMissingArgument
	EMissingArgument = "MissingArgument"
	// EMissingContentLength is oss's EMissingContentLength
	EMissingContentLength = "MissingContentLength"
	// ENoSuchBucket is oss's ENoSuchBucket
	ENoSuchBucket = "NoSuchBucket"
	// ENoSuchKey is oss's ENoSuchKey
	ENoSuchKey = "NoSuchKey"
	// ENoSuchUpload is oss's ENoSuchUpload
	ENoSuchUpload = "NoSuchUpload"
	// ENotImplemented is oss's ENotImplemented
	ENotImplemented = "NotImplemented"
	//EPreconditionFailed is oss's EPreconditionFailed
	EPreconditionFailed = "PreconditionFailed"
	//ERequestTimeTooSkewed is oss's ERequestTimeTooSkewed
	ERequestTimeTooSkewed = "RequestTimeTooSkewed"
	//ERequestTimeout is oss's ERequestTimeout
	ERequestTimeout = "RequestTimeout"
	// ERequestIsNotMultiPartContent is oss's ERequestIsNotMultiPartContent
	ERequestIsNotMultiPartContent = "RequestIsNotMultiPartContent"
	// ESignatureDoesNotMatch is oss's ESignatureDoesNotMatch
	ESignatureDoesNotMatch = "SignatureDoesNotMatch"
	// ETooManyBuckets is oss's ETooManyBuckets
	ETooManyBuckets = "TooManyBuckets"
)

var (
	// OSSAPIError is OSSAPI SDK's Error
	OSSAPIError = &Error{ErrNo: EAPI,
		ErrMsg:       "OSSAPI SDK's Inner Error,You Can Find More Details In Log Files",
		ErrDetailMsg: "OSSAPI SDK's Inner Error,You Can Find More Details In Log Files"}
	// ArgError is args' error such as nil
	ArgError = &Error{ErrNo: EAPI,
		ErrMsg:       "Argument Error.Please Check Your Argment Such as nil",
		ErrDetailMsg: "Argument Error.Please Check Your Argment Such as nil"}
)

// Error is ossapi's Error type
type Error struct {
	// XMLName is XML need
	XMLName xml.Name `xml:"Error"`
	//ErrNo is error code
	ErrNo string `xml:"Code"`
	// ErrMsg is brief error
	ErrMsg string `xml:"Message"`
	// HTTPStatus is http response
	HTTPStatus int
	// ErrDetailMsg is Detail message
	ErrDetailMsg string
}

func (e Error) Error() string {
	return e.ErrMsg
}
