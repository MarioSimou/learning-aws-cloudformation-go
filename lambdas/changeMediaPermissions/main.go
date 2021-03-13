package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func handler(ctx context.Context, event events.S3Event) error {

	var record = event.Records[0]
	var bucketName = record.S3.Bucket.Name
	var objectKey = record.S3.Object.Key
	var cfg, e = config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if e != nil {
		return e
	}

	var s3Client = s3.NewFromConfig(cfg)
	var aclInput = &s3.PutObjectAclInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		ACL:    types.ObjectCannedACLPublicRead,
	}
	if _, e := s3Client.PutObjectAcl(ctx, aclInput); e != nil {
		return e
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
