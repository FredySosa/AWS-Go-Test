package main

import (
	"github.com/FredySosa/AWS-Go-Test/createData/internal/container"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambdaHandler := container.Initialize()
	lambda.Start(lambdaHandler.LambdaHandler)
}
