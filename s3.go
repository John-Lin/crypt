package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/olekukonko/tablewriter"
)

func ListBucketObject(bucket, region string) {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(region)})
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	resp, err := svc.ListObjectsV2(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}
	data := make([][]string, 5)
	// fmt.Println(resp)
	for i, key := range resp.Contents {

		data = append(data, []string{strconv.Itoa(i + 1), *key.Key, strconv.FormatInt(*key.Size, 10)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No", "Key(Filename)", "Size(Bytes)"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func UploadSecret(bucket, region, localFileName, remoteFilename string) error {

	s3Uploader := s3manager.NewUploader(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	reader, err := os.Open(localFileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer reader.Close()

	input := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(remoteFilename),
		Body:   reader,
	}
	_, err = s3Uploader.Upload(input)
	if err != nil {
		return err
	}

	err = os.Remove(localFileName)
	if err != nil {
		return err
	}

	return nil
}

func DownloadSecret(bucket, region, remoteFilename string) (string, error) {
	s3Downloader := s3manager.NewDownloader(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	localFilename := path.Base(remoteFilename)
	f, err := os.Create(localFilename)
	if err != nil {
		return "", err
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(remoteFilename),
	}
	_, err = s3Downloader.Download(f, input)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
