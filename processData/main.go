package main

import (
	"context"
	"os"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/container"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	region := os.Getenv("AWS_REGION")
	ctx := context.Background()

	lambdaHandler := container.Initialize(ctx, region)
	lambda.Start(lambdaHandler.LambdaHandler)
}
