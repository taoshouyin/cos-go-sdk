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
		fmt.Println("Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("Code: %v\n", e.Code)
		fmt.Printf("Message: %v\n", e.Message)
		fmt.Printf("Resource: %v\n", e.Resource)
		fmt.Printf("RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Println(err)
		// ERROR
	}
}

func main() {
	u, _ := url.Parse("https://test-1259654469.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{
		BucketURL: u,
	}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	opt := &cos.BucketPutRefererOptions{
		Status:      "Enabled",
		RefererType: "White-List",
		DomainList: []string{
			"*.qq.com",
			"*.qcloud.com",
		},
		EmptyReferConfiguration: "Allow",
	}

	_, err := c.Bucket.PutReferer(context.Background(), opt)
	log_status(err)

	res, _, err := c.Bucket.GetReferer(context.Background())
	log_status(err)
	fmt.Printf("%+v\n", res)
}
