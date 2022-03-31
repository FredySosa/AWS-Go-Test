package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostsService_CreatePost(t *testing.T) {
	var (
		fakeRequest = domain.CreationRequest{
			URLToPost: "someurl.com",
			Text:      "Check this out!",
		}
		errGeneric = fmt.Errorf("some error")
	)
	type fields struct {
		PostsRepository func(m *mocks.MockPostsRepositoryPort)
	}
	type args struct {
		request domain.CreationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Response
		wantErr bool
	}{
		{
			name: "fail: empty url",
			args: args{
				request: domain.CreationRequest{
					URLToPost: "",
				},
			},
			want:    domain.Response{},
			wantErr: true,
		},
		{
			name: "fail: empty text",
			args: args{
				request: domain.CreationRequest{
					URLToPost: "Some url",
					Text:      "",
				},
			},
			want:    domain.Response{},
			wantErr: true,
		},
		{
			name: "fail: PostsRepository.CreatePost returns an error",
			fields: fields{
				PostsRepository: func(m *mocks.MockPostsRepositoryPort) {
					m.EXPECT().CreatePost(gomock.Any(), gomock.AssignableToTypeOf(domain.PostToSave{})).
						Return(errGeneric)
				},
			},
			args: args{
				request: fakeRequest,
			},
			want:    domain.Response{},
			wantErr: true,
		},
		{
			name: "success: full right test",
			fields: fields{
				PostsRepository: func(m *mocks.MockPostsRepositoryPort) {
					m.EXPECT().CreatePost(gomock.Any(), gomock.AssignableToTypeOf(domain.PostToSave{})).
						Return(nil)
				},
			},
			args: args{
				request: fakeRequest,
			},
			want: domain.Response{
				ID: "someID",
			},
			wantErr: false,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostsRepository := mocks.NewMockPostsRepositoryPort(mockCtrl)
			if tt.fields.PostsRepository != nil {
				tt.fields.PostsRepository(mockPostsRepository)
			}

			ps := PostsService{
				PostsRepository: mockPostsRepository,
			}
			got, err := ps.CreatePost(context.Background(), tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != "" {
				tt.want.ID = got.ID
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
