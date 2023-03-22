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

func setHotLink() {
	u, _ := url.Parse("https://test-1234567890.pic.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{CIURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	opt := &cos.HotLinkOptions{
		Url:  []string{"www.demo.com", "www.demo1.com", "www.demo2.com"},
		Type: "white",
	}

	_, err := c.CI.SetHotLink(context.Background(), opt)
	log_status(err)
}

func getHotLink() {
	u, _ := url.Parse("https://test-1234567890.pic.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{CIURL: u}
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

	res, _, err := c.CI.GetHotLink(context.Background())
	log_status(err)
	fmt.Printf("%+v\n", res)
}

func deleteHotLink() {
	u, _ := url.Parse("https://test-1234567890.pic.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{CIURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	opt := &cos.HotLinkOptions{
		Type: "off",
	}

	_, err := c.CI.SetHotLink(context.Background(), opt)
	log_status(err)
}

func main() {
	//setHotLink()
	//deleteHotLink()
	//getHotLink()
}
