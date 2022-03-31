package adapters

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/adapters/mocks"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
	"github.com/golang/mock/gomock"
)

func TestPostsRepository_CreatePost(t *testing.T) {
	var (
		fakePostToSave = domain.PostToSave{
			ID:        "someID",
			URLToPost: "someURL.com",
			Text:      "some test",
		}
		genericErr = fmt.Errorf("some error")
	)
	type fields struct {
		client func(m *mocks.MockDynamoDB)
	}
	type args struct {
		post domain.PostToSave
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "fail: client.PutItem returns an error",
			fields: fields{
				client: func(m *mocks.MockDynamoDB) {
					m.EXPECT().PutItem(gomock.Any(), gomock.AssignableToTypeOf(&dynamodb.PutItemInput{})).
						Return(nil, genericErr)
				},
			},
			args: args{
				post: fakePostToSave,
			},
			wantErr: true,
		},
		{
			name: "success: full right test",
			fields: fields{
				client: func(m *mocks.MockDynamoDB) {
					m.EXPECT().PutItem(gomock.Any(), gomock.AssignableToTypeOf(&dynamodb.PutItemInput{})).
						Return(&dynamodb.PutItemOutput{}, nil)
				},
			},
			args: args{
				post: fakePostToSave,
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

			mockClient := mocks.NewMockDynamoDB(mockCtrl)
			if tt.fields.client != nil {
				tt.fields.client(mockClient)
			}

			pr := PostsRepository{
				client: mockClient,
			}
			if err := pr.CreatePost(context.Background(), tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
