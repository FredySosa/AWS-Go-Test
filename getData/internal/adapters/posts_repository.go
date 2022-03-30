package adapters

import (
	"context"
	"fmt"

	"github.com/FredySosa/AWS-Go-Test/getData/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDB interface {
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
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

func (pr PostsRepository) GetPosts(ctx context.Context) ([]domain.Post, error) {
	out, err := pr.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("posts"),
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling results: %w", err)
	}
	if out == nil || len(out.Items) == 0 {
		return []domain.Post{}, nil
	}
	fmt.Println(out.Items)
	posts := make([]domain.Post, 0)
	err = attributevalue.UnmarshalListOfMaps(out.Items, &posts)
	if err != nil {
		return []domain.Post{}, nil
	}
	return posts, nil
}
