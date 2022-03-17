package main

import (
	"context"
	"os"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/container"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	region := os.Getenv("AWS_REGION")
	topic := "arn:aws:sns:us-east-1:526527389800:SendPosts"
	ctx := context.Background()

	lambdaHandler := container.Initialize(ctx, region, topic)
	lambda.Start(lambdaHandler.LambdaHandler)
}
