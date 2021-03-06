package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/ports"
	"github.com/aws/aws-lambda-go/events"
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

	var request domain.CreationRequest
	if err := json.NewDecoder(strings.NewReader(req.Body)).Decode(&request); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: domain.ParsingBodyError.HTTPCode,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: domain.ParsingBodyError.String(),
		}, nil
	}

	response, err := h.PostsServicePort.CreatePost(ctx, request)
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
