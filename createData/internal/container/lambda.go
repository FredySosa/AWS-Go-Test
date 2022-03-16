package container

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/adapters"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/services"
	"github.com/aws/aws-lambda-go/events"
)

type lambdaFunc func(
	ctx context.Context,
	request events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error)

type LambdaHandler struct {
	httpHandlerFunc lambdaFunc
}

func Initialize(ctx context.Context, region string) LambdaHandler {
	postsRepository := adapters.NewPostsRepository(ctx, region)
	postsService := services.NewPostsService(postsRepository)
	httpHandler := adapters.NewHTTPHandler(postsService)

	return LambdaHandler{
		httpHandlerFunc: httpHandler.ProcessRequest,
	}
}

func (lambda *LambdaHandler) LambdaHandler(
	ctx context.Context, req events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	return lambda.httpHandlerFunc(ctx, req)
}
