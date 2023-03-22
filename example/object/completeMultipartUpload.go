package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"

	"net/http"

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

func initUpload(c *cos.Client, name string) *cos.InitiateMultipartUploadResult {
	v, _, err := c.Object.InitiateMultipartUpload(context.Background(), name, nil)
	log_status(err)
	fmt.Printf("%#v\n", v)
	return v
}

func uploadPart(c *cos.Client, name string, uploadID string, blockSize, n int) string {

	b := make([]byte, blockSize)
	if _, err := rand.Read(b); err != nil {
		log_status(err)
	}
	s := fmt.Sprintf("%X", b)
	f := strings.NewReader(s)

	resp, err := c.Object.UploadPart(
		context.Background(), name, uploadID, n, f, nil,
	)
	log_status(err)
	fmt.Printf("%s\n", resp.Status)
	return resp.Header.Get("Etag")
}

func main() {
	u, _ := url.Parse("https://test-1259654469.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test/test_complete_upload.go"
	up := initUpload(c, name)
	uploadID := up.UploadID
	blockSize := 1024 * 1024 * 3

	opt := &cos.CompleteMultipartUploadOptions{}
	for i := 1; i < 5; i++ {
		etag := uploadPart(c, name, uploadID, blockSize, i)
		opt.Parts = append(opt.Parts, cos.Object{
			PartNumber: i, ETag: etag},
		)
	}

	c = cos.NewClient(b, &http.Client{
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
	v, resp, err := c.Object.CompleteMultipartUpload(
		context.Background(), name, uploadID, opt,
	)
	log_status(err)
	fmt.Printf("%s\n", resp.Status)
	fmt.Printf("%#v\n", v)
	fmt.Printf("%s\n", v.Location)
}
