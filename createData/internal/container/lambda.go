package container

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/adapters"

	"github.com/aws/aws-lambda-go/events"
)

type lambdaFunc func(
	ctx context.Context,
	request events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error)

type LambdaHandler struct {
	httpHandlerFunc lambdaFunc
}

func Initialize() LambdaHandler {
	httpHandler := adapters.NewHTTPHandler()
	return LambdaHandler{
		httpHandlerFunc: httpHandler.ProcessRequest,
	}
}

func (lambda *LambdaHandler) LambdaHandler(
	ctx context.Context, req events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	return lambda.httpHandlerFunc(ctx, req)
}
