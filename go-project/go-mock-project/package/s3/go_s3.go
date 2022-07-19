package go_s3

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func initS3Client() *s3.Client {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	return s3.NewFromConfig(cfg)
}

func S3PutObject(url string, v any) {

	client := initS3Client()

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("go-mock-project"),
		Key:    aws.String(createS3Path(url)),
		Body:   strings.NewReader(castToJsonString(v)),
	})
	if err != nil {
		panic(err)
	}
	// // Get the first page of results for ListObjectsV2 for a bucket
	// output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	// 	Bucket: aws.String("go-mock-project"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("first page results:")
	// for _, object := range output.Contents {
	// 	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	// }

}

func createS3Path(url string) string {
	now := time.Now()
	// Create path ex: fptshop.com.vn/YYYYMMDD.json
	return strings.Split(strings.Split(url, "//")[1], "/")[0] + "/" + now.Format("20060201") + ".json"

}

func castToJsonString(v any) string {
	dataJson, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(dataJson)
}
