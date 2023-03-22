package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/taoshouyin/cos-go-sdk"
	"github.com/taoshouyin/cos-go-sdk/debug"
)

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}

func describeMediaBucket() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	// DescirbeMediaBuckets 需要设置 CIURL 为 ci.<Region>.myqcloud.com
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	opt := &cos.DescribeMediaProcessBucketsOptions{
		Regions: "ap-chongqing",
	}
	res, _, err := c.CI.DescribeMediaProcessBuckets(context.Background(), opt)
	log_status(err)
	fmt.Printf("res: %+v\n", res)
}

func describePicBucket() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	// DescirbeMediaBuckets 需要设置 CIURL 为 ci.<Region>.myqcloud.com
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	opt := &cos.DescribePicProcessBucketsOptions{
		Regions: "ap-chongqing",
	}
	res, _, err := c.CI.DescribePicProcessBuckets(context.Background(), opt)
	log_status(err)
	fmt.Printf("res: %+v\n", res)
}

func describeAIBucket() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	// DescirbeMediaBuckets 需要设置 CIURL 为 ci.<Region>.myqcloud.com
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	opt := &cos.DescribeAIProcessBucketsOptions{
		Regions: "ap-chongqing",
	}
	res, _, err := c.CI.DescribeAIProcessBuckets(context.Background(), opt)
	log_status(err)
	fmt.Printf("res: %+v\n", res)
}

func describeASRBucket() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	// DescirbeMediaBuckets 需要设置 CIURL 为 ci.<Region>.myqcloud.com
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	opt := &cos.DescribeASRProcessBucketsOptions{
		Regions: "ap-chongqing",
	}
	res, _, err := c.CI.DescribeASRProcessBuckets(context.Background(), opt)
	log_status(err)
	fmt.Printf("res: %+v\n", res)
}

func main() {
	describeMediaBucket()
	describePicBucket()
	describeAIBucket()
	describeASRBucket()
}
