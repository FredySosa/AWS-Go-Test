package adapters

import (
	"context"
	"fmt"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const tableName = "posts"

//go:generate mockgen -destination=./mocks/dynamodb_mock.go -package=mocks -source=posts_repository.go
type DynamoDB interface {
	PutItem(
		ctx context.Context,
		params *dynamodb.PutItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.PutItemOutput, error)
}

type PostsRepository struct {
	client DynamoDB
}

func NewPostsRepository(ctx context.Context, region string) PostsRepository {
	configuration, _ := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = region

		return nil
	})

	return PostsRepository{
		client: dynamodb.NewFromConfig(configuration),
	}
}

func (pr PostsRepository) CreatePost(ctx context.Context, post domain.PostToSave) error {
	contactFields, err := attributevalue.MarshalMap(post)
	if err != nil {
		return fmt.Errorf("error marshaling contact: %w", err)
	}

	if _, err := pr.client.PutItem(
		ctx,
		&dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      contactFields,
		},
	); err != nil {
		return fmt.Errorf("error inserting data into db: %w", err)
	}

	return nil
}
