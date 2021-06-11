package config

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

// GetS3Session gets a reusable session to fetch data from s3
func GetS3Session(region string) *session.Session {
	if region == "" {
		log.Println("no region defined. using \"us-east-2\"")
		region = "us-east-2"
	}

	configP, _ := session.NewSession(&aws.Config{Region: aws.String(region)})

	return configP
}

func GetRegion(bucket string) string {
	sess := session.Must(session.NewSession())
	region, err := s3manager.GetBucketRegion(context.Background(), sess, bucket, "us-west-2")
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			fmt.Fprintf(os.Stderr, "unable to find bucket %s's region. bucket not found\n", bucket)
		}
		log.Fatal(err)
		return ""
	}

	return region
}
