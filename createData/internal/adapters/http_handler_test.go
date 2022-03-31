package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/mocks"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/ports"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ProcessRequest(t *testing.T) {
	var (
		fakeRequest = domain.CreationRequest{
			URLToPost: "someURL.com",
			Text:      "someText",
		}
		fakeBody, _ = json.Marshal(fakeRequest)
		genericErr  = fmt.Errorf("some error")
	)
	type fields struct {
		PostsServicePort func(m *mocks.MockPostsServicePort)
	}
	type args struct {
		req events.APIGatewayV2HTTPRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayV2HTTPResponse
		wantErr bool
	}{
		{
			name: "fail: cannot parse body",
			args: args{
				req: events.APIGatewayV2HTTPRequest{
					Body: "badBody",
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: domain.ParsingBodyError.HTTPCode,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: domain.ParsingBodyError.String(),
			},
			wantErr: false,
		},
		{
			name: "fail: PostsServicePort.CreatePost returns a generic error",
			fields: fields{
				PostsServicePort: func(m *mocks.MockPostsServicePort) {
					m.EXPECT().CreatePost(gomock.Any(), fakeRequest).
						Return(domain.Response{}, genericErr)
				},
			},
			args: args{
				req: events.APIGatewayV2HTTPRequest{
					Body: string(fakeBody),
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: domain.UnknownErr.HTTPCode,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: domain.UnknownErr.String(),
			},
			wantErr: false,
		},
		{
			name: "fail: PostsServicePort.CreatePost returns a custom error",
			fields: fields{
				PostsServicePort: func(m *mocks.MockPostsServicePort) {
					m.EXPECT().CreatePost(gomock.Any(), fakeRequest).
						Return(domain.Response{}, domain.CustomError{
							HTTPCode:     http.StatusTeapot,
							ErrorCode:    -1,
							MessageError: "something weird",
						})
				},
			},
			args: args{
				req: events.APIGatewayV2HTTPRequest{
					Body: string(fakeBody),
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: http.StatusTeapot,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"errorCode":-1,"messageError":"something weird"}`,
			},
			wantErr: false,
		},
		{
			name: "success: ful right test",
			fields: fields{
				PostsServicePort: func(m *mocks.MockPostsServicePort) {
					m.EXPECT().CreatePost(gomock.Any(), fakeRequest).
						Return(domain.Response{
							ID: "someID",
						}, nil)
				},
			},
			args: args{
				req: events.APIGatewayV2HTTPRequest{
					Body: string(fakeBody),
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: http.StatusCreated,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"id":"someID"}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostsService := mocks.NewMockPostsServicePort(mockCtrl)
			if tt.fields.PostsServicePort != nil {
				tt.fields.PostsServicePort(mockPostsService)
			}
			h := Handler{
				PostsServicePort: mockPostsService,
			}
			got, err := h.ProcessRequest(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewHTTPHandler(t *testing.T) {
	type args struct {
		sp ports.PostsServicePort
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPHandler(tt.args.sp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
