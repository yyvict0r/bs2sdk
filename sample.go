package main

import (
	"fmt"
	"os"
	"yy.com/bs2/bs2sdk"
)

func main() {
	Bucket := "hzttest"
	Endpoint := "hzttest.bs2ul.yy.com"
	AccessKeyId := "ak_imp"
	AccessKeySecret := "a350fa50a12886a4368e6308ec96521e97a62837"
	object_name := "test.txt"
	local_file_name := "test.txt"

	fmt.Println("test upload file [%s].", object_name)
	bs2client := bs2sdk.New(Bucket, Endpoint, AccessKeyId, AccessKeySecret)
	_, err := bs2client.Upload(object_name, local_file_name, "")
	if err != nil {
		// handle error.
	}

	fmt.Println("test download file [%s].", object_name)
	bs2client = bs2sdk.New(Bucket, Endpoint, AccessKeyId, AccessKeySecret)
	//_, err = bs2client.MultipartDownload(object_name, local_file_name, 5)
	_, err = bs2client.Download(object_name, local_file_name)
	if err != nil {
		// handle error.
	}

	fmt.Println("test delete file [%s].", object_name)
	bs2client = bs2sdk.New(Bucket, Endpoint, AccessKeyId, AccessKeySecret)
	_, err = bs2client.Delete(object_name)
	if err != nil {
		// handle error.
	}

	os.Exit(0)
}
