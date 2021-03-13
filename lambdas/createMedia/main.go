package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
)

var generic720pPresetId = "1351620000001-000010"
var generic360pPresetId = "1351620000001-000061"
var pipelineID = "1615643457138-1ciklk"

func handler(ctx context.Context, event events.S3Event) error {
	var record = event.Records[0]
	var objectName = record.S3.Object.Key
	var ext = filepath.Ext(objectName)
	var name = strings.Replace(objectName, ext, "", -1)
	var cfg, e = config.LoadDefaultConfig(context.Background(), config.WithRegion("eu-west-1"))
	if e != nil {
		return e
	}

	var elasticTranscoderClient = elastictranscoder.NewFromConfig(cfg)

	var jobInput = elastictranscoder.CreateJobInput{
		PipelineId: aws.String(pipelineID),
		Input: &types.JobInput{
			Key: aws.String(objectName),
		},
		Outputs: []types.CreateJobOutput{
			types.CreateJobOutput{
				Key:      aws.String(fmt.Sprintf("%s-720%s", name, ext)),
				PresetId: aws.String(generic720pPresetId),
			},
			types.CreateJobOutput{
				Key:      aws.String(fmt.Sprintf("%s-360%s", name, ext)),
				PresetId: aws.String(generic360pPresetId),
			},
		},
	}

	if _, e := elasticTranscoderClient.CreateJob(context.Background(), &jobInput); e != nil {
		return e
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
