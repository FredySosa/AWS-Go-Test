package container

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/adapters"
	"github.com/FredySosa/AWS-Go-Test/processData/internal/core/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type lambdaFunc func(ctx context.Context, e events.DynamoDBEvent) error

type LambdaHandler struct {
	httpHandlerFunc lambdaFunc
}

func Initialize(ctx context.Context, region, topic string) LambdaHandler {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	}))

	svc := sns.New(sess)
	snsService := services.NewSNSService(svc, topic)
	snsHandler := adapters.NewSNSHandler(snsService)

	return LambdaHandler{
		httpHandlerFunc: snsHandler.ProcessRequest,
	}
}

func (lambda *LambdaHandler) LambdaHandler(ctx context.Context, e events.DynamoDBEvent) error {

	return lambda.httpHandlerFunc(ctx, e)
}
