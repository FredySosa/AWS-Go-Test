package adapters

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
}

func NewHTTPHandler() Handler {
	return Handler{}
}

func (h Handler) ProcessRequest(
	ctx context.Context,
	request events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	body := request.Body
	response := "##"
	if strings.Contains(body, "test") {
		response = "@@"
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: fmt.Sprintf(`{"response":"%s"}`, response),
	}, nil
}
