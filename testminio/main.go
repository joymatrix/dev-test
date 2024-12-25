package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 忽略证书验证
		},
	}

	endPoint := fmt.Sprintf("%s:%d", "47.112.115.51", 9000)

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:     credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure:    true,
		Transport: transport,
	})
	if err != nil {
		log.Fatal("minio client create fail, err %+v", err)
	}
	fmt.Println("connect success")
	reqParams := url.Values{}
	url, err := minioClient.PresignedGetObject(context.Background(), "relax-park-311", "1.jpeg", time.Duration(10000)*time.Second, reqParams)
	if err != nil {
		fmt.Printf("get object url err:%+v", err.Error())
		return
	}
	fmt.Println("url:", url.String())

}
