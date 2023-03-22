package main

import (
	"context"
	"net/url"
	"os"

	"net/http"

	"github.com/taoshouyin/cos-go-sdk"
	"github.com/taoshouyin/cos-go-sdk/debug"
)

func main() {
	u, _ := url.Parse("https://alanbj-1251668577.cos.ap-beijing.myqcloud.com")
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

	opt := &cos.BucketPutVersionOptions{
		// Enabled or Suspended, the versioning once opened can not close.
		Status: "Enabled",
	}

	_, err := c.Bucket.PutVersioning(context.Background(), opt)
	if err != nil {
		panic(err)
	}
}
