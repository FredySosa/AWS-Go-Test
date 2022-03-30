package adapters

import (
	"context"
	"errors"
	"github.com/FredySosa/AWS-Go-Test/getData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/getData/internal/ports"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Handler struct {
	PostsServicePort ports.PostsServicePort
}

func NewHTTPHandler(sp ports.PostsServicePort) Handler {
	return Handler{
		PostsServicePort: sp,
	}
}

func (h Handler) ProcessRequest(
	ctx context.Context,
	req events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	response, err := h.PostsServicePort.GetPosts(ctx)
	if err != nil {
		toReturn := events.APIGatewayV2HTTPResponse{
			StatusCode: domain.UnknownErr.HTTPCode,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: domain.UnknownErr.String(),
		}
		var ce domain.CustomError
		if errors.As(err, &ce) {
			toReturn.StatusCode = ce.HTTPCode
			toReturn.Body = ce.String()
		}

		return toReturn, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: response.String(),
	}, nil
}
