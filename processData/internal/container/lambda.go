package container

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/adapters"
	"github.com/aws/aws-lambda-go/events"
)

type lambdaFunc func(ctx context.Context, e events.DynamoDBEvent) error

type LambdaHandler struct {
	httpHandlerFunc lambdaFunc
}

func Initialize(ctx context.Context, region string) LambdaHandler {
	snsHandler := adapters.NewSNSHandler()

	return LambdaHandler{
		httpHandlerFunc: snsHandler.ProcessRequest,
	}
}

func (lambda *LambdaHandler) LambdaHandler(ctx context.Context, e events.DynamoDBEvent) error {

	return lambda.httpHandlerFunc(ctx, e)
}
