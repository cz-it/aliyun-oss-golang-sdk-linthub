/**
* Author: CZ cz.theng@gmail.com
 */
//Service package wraps action for a user. Such as Query all your buckets (And Now  Have The Only Fuction).
//
//##service.QueryBuckets
//
//	QueryBuckets(prefix, marker string, maxKeys int) (bucketsInfo *BucketsInfo, ossapiError *ossapi.Error)
//
//QueryBuckets will Query all your buckets.
//
//* if prefix provided, buckets returned only have such prefix
//* if marker provided, buckets returned only after this marker
//* if maxKeys provided, at most maxKeys items returned
//
//
//BucketsInfo holds the bucket's meta information.
//
//	type Owner struct {
//	    ID          string
//	    DisplayName string
//	}
//
//	type Bucket struct {
//	    Name         string
//	    CreationDate string
//	    Location     string
//	}
//
//	type BucketsInfo struct {
//	    XMLName     xml.Name `xml:"ListAllMyBucketsResult"`
//	    Prefix      string   `xml:"Prefix"`
//	    Marker      string   `xml:"Marker"`
//	    MaxKeys     int      `xml:"MaxKeys"`
//	    IsTruncated bool     `xml:"IsTruncated"`
//	    NextMarker  string   `xml:"NextMarker"`
//	    Owner       Owner    `xml:"Owner"`
//	    Buckets     Buckets  `xml:"Buckets"`
//	}
//
//meta information such as Marker/Prefix/Owner is in BucketsInfo. And Buckets holds the real one .It is a list of Bucket recoding every bucket's name/location and create time.
//

package service

import ()
